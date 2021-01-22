package service

import (
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
	"github.com/teten-nugraha/golang-twitter-clone-api/mapper"
	"github.com/teten-nugraha/golang-twitter-clone-api/repository"
)

type UserServiceContract interface {

	/*
	 * Mencari User by username equal
	 */
	FindByUsername(username string) (dto.UserDto,error)

}

type UserService struct {
	UserRepository repository.UserRepository
}

func ProviderUserService(u repository.UserRepository) UserService {
	return UserService{
		UserRepository: u,
	}
}

func (u *UserService) FindByUsername(username string) (dto.UserDto, error) {
	user, err := u.UserRepository.FindByUsername(username)
	if err != nil {
		return mapper.ToUserDto(user), err
	}
	return mapper.ToUserDto(user), nil
}
