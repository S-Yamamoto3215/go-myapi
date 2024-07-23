package services

import (
	"github.com/S-Yamamoto3215/go-myapi/models"
	"github.com/S-Yamamoto3215/go-myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
