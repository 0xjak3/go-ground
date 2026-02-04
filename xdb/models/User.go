package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role" gorm:"not null;default:'user'"`
}
