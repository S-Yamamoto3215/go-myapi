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

// TODO: 指定IDの記事についたコメント一覧を取得。取得したコメントデータと発生したエラーを返す
func SelectCommentList() ([]models.Comment, error) {
	return []models.Comment{}, nil
}
