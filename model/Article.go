package model

import (
	"ginblog/utils/errmsg"
	"log"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid;references:ID"`
	Title    string   `gorm:"type:varchar(128);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(1024)" json:"desc"`
	Content  string   `gorm:"type:text" json:"content"`
	Img      string   `gorm:"type:varchar(128)" json:"img"`
}

// CreateArticle
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// search all articles with one category
func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	// var art Article
	var cateArtList []Article
	var total int64
	// db.Model(&art).Joins("inner join categories on articles.cid = ?", cid).Limit(pageSize).Offset((pageNum - 1) * pageSize).Scan(&cateArtList)
	// db.Model(&art).Where("cid = ?", cid).Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Scan(&cateArtList)
	// db.Model(&art).Preload("Category").Select("*").Where("cid = ?", id).Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("left join categories on articles.cid = categories.id").Scan(&cateArtList)

	// err := db.Preload("Categories.cid").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&cateArtList, "cid = ?", id).Error
	// db.Model(&cateArtList).Count(&total)

	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", cid).Find(&cateArtList).Error
	db.Model(&cateArtList).Where("cid =?", cid).Count(&total)

	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// search one article
func GetArtInfo(id int) (Article, int) {
	var art Article
	err = db.Where("id = ?", id).Preload("Category").First(&art).Error
	// db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// GetArticle list
func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var art Article
	var artList []Article
	var err error
	var total int64

	// err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageNum).Error
	// err := db.Find(&artList).Limit(pageSize).Offset((pageNum - 1) * pageNum).Error
	// db.Model(&artList).Count(&total)
	// err = db.Select("articles.id, title, img, created_at, updated_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("categories").Find(&artList).Error
	// err = db.Model(&art).Joins("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Scan(&artList).Error
	err = db.Model(&art).Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("Category").Find(&artList).Error
	log.Println("artList: ", artList)
	db.Model(&artList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return artList, errmsg.SUCCESS, total
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
