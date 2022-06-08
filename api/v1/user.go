package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

// check if user exist
func UserExist() {
	//
}

// add user
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// search single user

// search user list
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize < 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}

	// if pageSize == 0 {
	// 	pageSize = -1
	// }

	// if pageNum == 0 {
	// 	pageNum = -1
	// }
	code = errmsg.SUCCESS
	data, total := model.GetUsers(username, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// upate user
func EditUser(c *gin.Context) {

}

// delete user
func DeleteUser(c *gin.Context) {

}
