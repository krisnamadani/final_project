package models

type Photo struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	Photo_URL  string `json:"photo_url"`
	User_ID    uint   `json:"user_id"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
	User       User
}
