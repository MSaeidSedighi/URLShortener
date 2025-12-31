package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model

	OriginalUrl string
	ShortCode   string `gorm:"uniqueIndex"`
	Visits      uint64

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
