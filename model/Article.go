package model

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(128);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(1024)" json:"desc"`
	Content string `gorm:"type:text" json:"content"`
	Img     string `gorm:"type:varchar(128)" json:"img"`
}
