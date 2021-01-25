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
		Tweet: entity.Tweet,
		Maker: entity.UserRefer,
		LikeCount: entity.LikeCount,
		CommentCount: entity.CommentCount,
		CreatedAt: entity.CreatedAt,
	}
}

func EntityToTweetDtoList(entities []domain.Tweet) []dto.TweetDto {
	tweetDtoList := make([]dto.TweetDto, len(entities))

	for i, itm := range  entities {
		tweetDtoList[i] = EntityToTweetDto(itm)
	}

	return tweetDtoList
}
