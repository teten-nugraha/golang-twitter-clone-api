package mapper

import (
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
)

func NewTweetDtoToEntity(newTweet dto.NewTweetDto) domain.Tweet {
	return domain.Tweet{
		Tweet: newTweet.Tweet,
		UserRefer: newTweet.Maker,
	}
}

func EntityToTweetDto(entity domain.Tweet) dto.TweetDto {
	return dto.TweetDto{
		ID: entity.ID,
		Maker: "Test",
		CreatedAt: entity.CreatedAt,
	}
}
