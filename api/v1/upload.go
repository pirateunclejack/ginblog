package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("Get file from formfile error: ", err)
	}
	fileSize := fileHeader.Size

	url, code := model.UploadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})

}
