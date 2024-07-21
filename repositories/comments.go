package repositories

import (
	"database/sql"

	"github.com/S-Yamamoto3215/go-myapi/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		INSERT INTO comments (article_id, message, created_at) VALUES
		(?, ?, NOW());
	`

	var newComment models.Comment
	newComment.AritcleID, newComment.Message = comment.AritcleID, comment.Message

	result, err := db.Exec(sqlStr, comment.AritcleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		SELECT *
		from comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		rows.Scan(&comment.CommentID, &comment.AritcleID, &comment.Message, &comment.CreatedAt)
		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
