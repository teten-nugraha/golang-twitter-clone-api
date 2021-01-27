package service

import (
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
	"github.com/teten-nugraha/golang-twitter-clone-api/mapper"
	"github.com/teten-nugraha/golang-twitter-clone-api/repository"
)

type CommentServiceContract interface {
	SaveComment(newCommentDto dto.CommentDto) (dto.CommentDto, error)
	Comments(tweetId uint64) []dto.CommentDto
}

type CommentService struct {
	CommentRepository repository.CommentRepository
	TweetRepository repository.TweetRepository
}

func ProviderCommentService(c repository.CommentRepository, t repository.TweetRepository) CommentService {
	return CommentService{
		CommentRepository: c,
		TweetRepository: t,
	}
}

// implementasi
func (c *CommentService) SaveComment(newCommentDto dto.CommentDto) (dto.CommentDto, error) {

	newComment := mapper.CommentDtoToEntity(newCommentDto)

	newComment, err := c.CommentRepository.SaveComment(newComment)
	if err != nil {
		return dto.CommentDto{}, err
	}

	tweet_id := newComment.TweetId
	err = c.TweetRepository.UpdateComment(tweet_id)
	if err != nil {
		return dto.CommentDto{}, nil
	}

	newCommentDto = mapper.CommentEntityToDto(newComment)

	return newCommentDto, nil

}

func (c *CommentService) Comments(tweetId uint64) []dto.CommentDto {

	comments := c.CommentRepository.Comments(tweetId)

	commentDtos := mapper.CommentEntityToDtoList(comments)

	return commentDtos
}