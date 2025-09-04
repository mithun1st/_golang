package model

type UserModel struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
