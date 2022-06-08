package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(128);not null" json:"name"`
}
