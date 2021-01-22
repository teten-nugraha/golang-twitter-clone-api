package mapper

import (
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
)

func ToUserDto(user domain.User) dto.UserDto {
	return dto.UserDto{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Gender: user.Gender,
	}
}