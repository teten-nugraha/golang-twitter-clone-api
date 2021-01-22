package dto

type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignupDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"role" validate:"required"`
	Email    string `json:"email" validate:"email"`
}
