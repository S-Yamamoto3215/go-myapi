package repositories

import (
	"database/sql"

	"github.com/S-Yamamoto3215/go-myapi/models"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	INSERT INTO articles (title, contents, username, nice, created_at) VALUES
	(?, ?, ?, 0, NOW())`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()
	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		SELECT article_id, title, contents, username, nice
		from articles
		limit ? ofset ?;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

// TODO: 投稿IDを指定して記事データを取得。取得した記事データと発生したエラーを返す
func SelectArticleDetail() (models.Article, error) {
	return models.Article{}, nil
}

// TODO: いいねの数をアップデートする。発生したエラーを返す
func UpdateNiceNum() error {
	return nil
}
