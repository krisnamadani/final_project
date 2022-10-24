package models

type Comment struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	User_ID    uint   `json:"user_id"`
	Photo_ID   uint   `json:"photo_id"`
	Message    string `json:"message"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
	Photo      Photo
	User       User
}
