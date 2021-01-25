package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
	"github.com/teten-nugraha/golang-twitter-clone-api/mapper"
	"github.com/teten-nugraha/golang-twitter-clone-api/repository"
	"github.com/teten-nugraha/golang-twitter-clone-api/util"
	"time"
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

	CheckLogin(username, password string) (domain.User, error)
	GenerateToken(user domain.User) (string, error)

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

func (u *UserService) CheckLogin(username, password string) (domain.User, error) {

	user, err := u.UserRepository.FindByUsername(username)

	if err != nil {
		return user, nil
	}

	match, err := util.CheckPasswordHash(password, user.Password)
	if !match {
		logrus.Error("Hash and Password doesnt match")
		return user, err
	}

	return user, nil
}

func (u *UserService) GenerateToken(user domain.User) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(util.SECRET))

	return util.BEARER + t, err

}