package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
)

type TweetRepositoryContract interface {
	SaveTweet(newTweet domain.Tweet) (domain.Tweet, error)
}

type TweetRepository struct {
	DB *gorm.DB
}

func ProviderTweetRepository(DB *gorm.DB) TweetRepository {
	return TweetRepository{
		DB: DB,
	}
}

// implementasi
func (t *TweetRepository) SaveTweet(newTweet domain.Tweet) (domain.Tweet, error) {
	if err := t.DB.Create(&newTweet).Error; err != nil {
		return newTweet, err
	}

	return newTweet, nil
}