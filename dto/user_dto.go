package dto

type UserDto struct {
	ID                 uint64              `json:"id,string,omitempty"`
	Username           string              `json:"username"`
	Email              string              `json:"email"`
	Gender             string              `json:"gender"`
}
