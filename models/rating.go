package models

type Rating struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Stars   string `json:"stars" gorm:"not null"`
	Comment string `json:"comment"`
	MediaId *int   `json:"media_id" gorm:"default:null"`
}
