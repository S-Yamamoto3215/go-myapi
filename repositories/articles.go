package repositories

import "github.com/S-Yamamoto3215/go-myapi/models"

// TODO: 新規投稿をDBにインサートする。保存内容と、発生したエラーを返す
func InsertArticle() (models.Article, error) {
	return models.Article{}, nil
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
