package repositories

import (
	"database/sql"

	"github.com/S-Yamamoto3215/go-myapi/models"
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

// TODO: 変数 page で指定されたページに表示する投稿一覧をデータベースから取得する関数。取得した記事データと発生したエラーを返す
func SelectArticleList() ([]models.Article, error) {
	return []models.Article{}, nil
}

// TODO: 投稿IDを指定して記事データを取得。取得した記事データと発生したエラーを返す
func SelectArticleDetail() (models.Article, error) {
	return models.Article{}, nil
}

// TODO: いいねの数をアップデートする。発生したエラーを返す
func UpdateNiceNum() error {
	return nil
}
