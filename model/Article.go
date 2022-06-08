package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(128);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(1024)" json:"desc"`
	Content string `gorm:"type:text" json:"content"`
	Img     string `gorm:"type:varchar(128)" json:"img"`
}

// CreateArticle
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo search all articles with one category

// todo search one article

// todo GetArticle list
func GetArt(title string, pageSize int, pageNum int) ([]Article, int64) {
	var art []Article
	var total int64

	err := db.Find(&art).Limit(pageSize).Offset((pageNum - 1) * pageNum).Error
	db.Model(&art).Count(&total)
	if err != nil {
		return nil, 0
	}
	return art, total
}

// EditArticle
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&art).Where("id = ?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArticle
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
