package models

type Category struct {
	Id           uint    `json:"id" gorm:"primaryKey"`
	CategoryName string  `json:"category_name" gorm:"unique;not null"`
	Desc         string  `json:"desc"`
	Media        []Media `gorm:"constraint:OnDelete:SET NULL;"`
}
