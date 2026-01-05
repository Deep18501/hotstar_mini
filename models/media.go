package models

type Media struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Title        string `json:"title" gorm:"not null"`
	Desc         string `json:"desc"`
	ThumbnailURL string `json:"thumbnail_url"`
	BannerURL    string `json:"banner_url"`
	MediaUrl     string `json:"media_url"`
	AgeRating    string `json:"age_rating"`
	ReleaseYear  string `json:"release_year"`
	CategoryId   *int   `json:"category_id" gorm:"default:null"`
	Rating       []*Rating
	Genre        []*Genre `gorm:"many2many:media_genres;"`
}
