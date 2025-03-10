package models

import "gorm.io/gorm"

type LongShort struct{
	LongUrl string `json:"url"`
	ShortUrl string `json:"shorturl" gorm:"unique"`
	gorm.Model
}