package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
)

type CommentRepositoryContract interface {
	SaveComment(newComment domain.Comment) (domain.Comment, error)
}

type CommentRepository struct {
	DB *gorm.DB
}

func ProviderCommentRepository(DB *gorm.DB) CommentRepository {
	return CommentRepository{
		DB: DB,
	}
}

// implementasi
func (c *CommentRepository) SaveComment(newComment domain.Comment) (domain.Comment, error) {
	if err := c.DB.Create(&newComment).Error; err != nil {
		logrus.Warn("Failed save comment")
		return newComment, err
	}

	logrus.Info("Success save new comment")
	return newComment, nil
}