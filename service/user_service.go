package service

import (
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
	"github.com/teten-nugraha/golang-twitter-clone-api/mapper"
	"github.com/teten-nugraha/golang-twitter-clone-api/repository"
	"github.com/teten-nugraha/golang-twitter-clone-api/util"
)

type UserServiceContract interface {

	/*
	 * Mencari User by username equal
	 */
	FindByUsername(username string) (dto.UserDto,error)

	/**
	 * Untuk signup user baru
	 */
	SignUp(dto dto.SignupDto) (dto.UserDto, error)

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

func(u *UserService) SignUp(signUpDto dto.SignupDto) (dto.UserDto, error) {

	var userDto dto.UserDto

	user := mapper.SignupDtoToEntity(signUpDto)

	hash, _ :=util.HashPassword(signUpDto.Password)
	user.Password = hash

	user, err := u.UserRepository.SaveOrUpdate(user)
	if err != nil {
		return userDto, err
	}

	userDto, err = mapper.ToUserDto(user), nil

	return userDto, err
}