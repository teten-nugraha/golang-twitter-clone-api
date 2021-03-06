package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
)

type TweetRepositoryContract interface {
	SaveTweet(newTweet domain.Tweet) (domain.Tweet, error)
	Timeline() []domain.Tweet
	UpdateComment(idTweet uint64) (error)
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
		logrus.Warn("failed save new tweet")
	}

	logrus.Info("success save new tweet")
	return newTweet, nil
}

func (t *TweetRepository) Timeline() []domain.Tweet {
	var tweets []domain.Tweet

	t.DB.Order("id desc").Find(&tweets)

	logrus.Info("Success get tweets")

	return tweets

}

func (t *TweetRepository) UpdateComment(idTweet uint64) (error) {
	var tweet domain.Tweet

	t.DB.Where("id = ?", idTweet).First(&tweet)

	tweet.CommentCount += 1

	t.DB.Model(&tweet).Where("id = ?", idTweet).Update("comment_count", tweet.CommentCount)

	return nil
}