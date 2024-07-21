package repositories

import "github.com/S-Yamamoto3215/go-myapi/models"

// TODO: 新規投稿をDBにインサートする。保存内容と、発生したエラーを返す
func InsertComment() (models.Comment, error) {
	return models.Comment{}, nil
}

// TODO: 指定IDの記事についたコメント一覧を取得。取得したコメントデータと発生したエラーを返す
func SelectCommentList() ([]models.Comment, error) {
	return []models.Comment{}, nil
}
