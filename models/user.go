package models

type User struct {
	ID int `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age int8 `json:"age" binding:"gte=1,lte=130"`
}