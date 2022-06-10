package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

// generate token
func GenToken(username string, password string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	expireAt := jwt.NewNumericDate(expireTime)
	SetClaims := MyClaims{
		Username: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expireAt,
			Issuer:    "ginblog",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodES256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS

}

// verify token
func CheckToken(token string) (*MyClaims, int) {
	settoken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := settoken.Claims.(*MyClaims); ok && settoken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// jwt middleware
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_WRONG_FORMAT
			c.Abort()
		}
		key, Tcode := CheckToken(checkToken[1])
		if Tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt.Unix() {
			code = errmsg.ERROR_TOKEN_EXPIRED
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		c.Set("Username", key.Username)
		c.Next()
	}
}
