package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(128);not null" json:"username"`
	Password string `gorm:"type:varchar(128);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// check if user exist
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED // 1001
	}
	return errmsg.SUCCESS
}

// add new user
func CreateUser(data *User) int {
	// data.Password = ScryptPassword(data.Password)
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// check user list
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	// err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	// offset := (pageNum - 1) * pageSize
	// if pageNum == -1 && pageSize == -1 {
	// 	offset = -1
	// }

	// err := db.Limit(pageSize).Offset(offset).Find(&users).Error

	if username != "" {
		db.Select("id,username,role,created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	// err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	if err != nil {
		return users, 0
	}
	return users, total
}

// update user
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role

	err := db.Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// delete user
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// scrypt
func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

func ScryptPw(password string) string {
	salt := []byte{12, 32, 4, 32, 95, 98, 85, 75}
	HashPw, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	FinalPw := base64.StdEncoding.EncodeToString(HashPw)
	return FinalPw
}
