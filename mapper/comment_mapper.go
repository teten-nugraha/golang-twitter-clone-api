package mapper

import (
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
)

func CommentEntityToDto(entity domain.Comment) dto.CommentDto {
	return dto.CommentDto{
		ID: entity.ID,
		Maker: entity.Maker,
		Reply: entity.Reply,
		CreatedAt: entity.CreatedAt,
		TweetId: entity.TweetId,
	}
}

func CommentDtoToEntity(dto dto.CommentDto) domain.Comment {
	return domain.Comment{
		Maker: dto.Maker,
		Reply: dto.Reply,
		TweetId: dto.TweetId,
	}
}

func CommentEntityToDtoList(entities []domain.Comment) []dto.CommentDto {
	commentDtoList := make([]dto.CommentDto, len(entities))

	for i, itm := range  entities {
		commentDtoList[i] = CommentEntityToDto(itm)
	}

	return commentDtoList
}