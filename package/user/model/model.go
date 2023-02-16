package model

type User struct {
	Id        string `json:"id" gorm:"column:id"`
	Username  string `json:"username" gorm:"column:username"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"last_name"`
}
