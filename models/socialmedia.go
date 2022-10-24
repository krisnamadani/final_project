package models

type SocialMedia struct {
	ID               uint   `json:"id" gorm:"primary_key"`
	Name             string `json:"name"`
	Social_Media_URL string `json:"social_media_url"`
	User_ID          uint   `json:"user_id"`
	Created_At       string `json:"created_at"`
	Updated_At       string `json:"updated_at"`
	User             User
}
