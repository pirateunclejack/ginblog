package model

import (
	"ginblog/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(128);not null" json:"name"`
}

// CheckCategory
func CheckCate(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// CreateCategory
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategory list
func GetCate(name string, pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64

	err := db.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageNum).Error
	db.Model(&cate).Count(&total)
	if err != nil {
		return nil, 0
	}
	return cate, total
}

// EditCategory
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// check all articles on this category

// DeleteCategory
func DeleteCate(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
