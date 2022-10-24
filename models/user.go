package models

type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Username   string `json:"username" gorm:"unique_index"`
	Email      string `json:"email" gorm:"unique_index"`
	Password   string `json:"password"`
	Age        uint   `json:"age"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
