package service

import (
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
	"github.com/teten-nugraha/golang-twitter-clone-api/mapper"
	"github.com/teten-nugraha/golang-twitter-clone-api/repository"
)

type TweetServiceContract interface {
	SaveTweet(newTweetDto dto.NewTweetDto, maker uint) (dto.TweetDto, error)
}

type TweetService struct {
	TweetRepository repository.TweetRepository
}

func ProviderTweetService(u repository.TweetRepository) TweetService {
	return TweetService{
		TweetRepository: u,
	}
}


// implementasi
func (t *TweetService) SaveTweet(newTweetDto dto.NewTweetDto) (dto.TweetDto, error) {

	var tweet domain.Tweet

	tweet.UserRefer = newTweetDto.Maker
	tweet.Tweet = newTweetDto.Tweet

	tweet, err := t.TweetRepository.SaveTweet(tweet)
	if err != nil {
		return dto.TweetDto{}, err
	}

	var tweetDto = mapper.EntityToTweetDto(tweet)

	return tweetDto, nil

}