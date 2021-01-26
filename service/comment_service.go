package service

import (
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
	"github.com/teten-nugraha/golang-twitter-clone-api/mapper"
	"github.com/teten-nugraha/golang-twitter-clone-api/repository"
)

type CommentServiceContract interface {
	SaveComment(newCommentDto dto.CommentDto) (dto.CommentDto, error)
}

type CommentService struct {
	CommentRepository repository.CommentRepository
}

func ProviderCommentService(c repository.CommentRepository) CommentService {
	return CommentService{
		CommentRepository: c,
	}
}

// implementasi
func (c *CommentService) SaveComment(newCommentDto dto.CommentDto) (dto.CommentDto, error) {

	newComment := mapper.CommentDtoToEntity(newCommentDto)

	newComment, err := c.CommentRepository.SaveComment(newComment)
	if err != nil {
		return dto.CommentDto{}, err
	}

	newCommentDto = mapper.CommentEntityToDto(newComment)

	return newCommentDto, nil

}