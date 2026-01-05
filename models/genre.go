package models

type Genre struct {
	Id        uint    `json:"id" gorm:"primaryKey"`
	GenreType string  `json:"genre" gorm:"unique;not null"`
	Media     []*Media `gorm:"many2many:media_genres;"`
}
