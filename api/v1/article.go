package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateArticle
func AddArt(c *gin.Context) {
	var data model.Article
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("AddArt bind json error: ", err)
	}

	code = model.CreateArt(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// search all articles with a category
func GetCateArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Param("cid"))
	data, code, total := model.GetCateArt(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"total":   total,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// search one article
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticle list
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize < 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}

	if len(title) == 0 {
		data, code, total := model.GetArt(pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	// data, code, total := model.SearchArt(title, pageSize, pageNum)
	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  code,
	// 	"data":    data,
	// 	"total":   total,
	// 	"message": errmsg.GetErrMsg(code),
	// })
}

// EditArticle
func EditArt(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.EditArt(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteArticle
func DeleteArt(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))

	code := model.DeleteArt(cid)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
