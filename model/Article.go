package model

import (
	"ginblog/utils/errmsg"
	"log"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid;regerences:ID"`
	gorm.Model
	Title    string   `gorm:"type:varchar(128);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Describe     string   `gorm:"type:varchar(1024)" json:"describe"`
	Content  string   `gorm:"type:text" json:"content"`
	Img      string   `gorm:"type:varchar(128)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// CreateArticle
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// search all article with one category
func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	// var art Article
	var cateArtList []Article
	var total int64
	// db.Model(&art).Joins("inner join category on article.cid = ?", cid).Limit(pageSize).Offset((pageNum - 1) * pageSize).Scan(&cateArtList)
	// db.Model(&art).Where("cid = ?", cid).Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Scan(&cateArtList)
	// db.Model(&art).Preload("Category").Select("*").Where("cid = ?", id).Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("left join category on article.cid = category.id").Scan(&cateArtList)

	// err := db.Preload("category.cid").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&cateArtList, "cid = ?", id).Error
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
	var articleList []Article
	var err error
	var total int64
	err = db.Preload("Category").
		Select("article.id, title, img, cid, created_at, updated_at, describe, comment_count, read_count, category.name").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("Created_At DESC").
		Joins("LEFT JOIN category ON article.cid = category.id").
		Find(&articleList).Error
	// count
	db.Model(&articleList).Count(&total)
	log.Println(articleList)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// EditArticle
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["describe"] = data.Describe
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

// Search article via title
func SearchArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = db.Preload("Category").
		Select("article.id, title, img, cid, created_at, updated_at, describe, comment_count, read_count, category.name").
		Order("Created_At DESC").Joins("LEFT JOIN category ON article.cid = category.id").
		Where("title LIKE ?", title+"%",).
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	// Count
	db.Model(&articleList).Where("title LIKE ?",
		title+"%",
	).Count(&total)
	log.Println(articleList)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}
