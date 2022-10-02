package dto

type RegisterDto struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}
