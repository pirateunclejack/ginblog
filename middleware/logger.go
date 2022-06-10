package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	logPath := "./log/ginblog.log"
	linkName := "./log/ginblog.latest.log"
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Open ginblog.log with error: ", err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		logPath+".%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(linkName),
	)
	if err != nil {
		fmt.Println("rotetelogs generate error: ", err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.DebugLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}

	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%v ms", stopTime.Nanoseconds()/1000000)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"Hostname":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"ClientIP":  clientIP,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Error()
		} else {
			entry.Info()
		}

	}
}
