package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... , user mode error
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USER_NOT_EXIST     = 1003
	ERROR_TOKEN_NOT_EXIST    = 1004
	ERROR_TOKEN_EXPIRED      = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_WRONG_FORMAT = 1007
	ERROR_USER_NO_RIGHT      = 1008
	// code = 2000... , article mode error
	ERROR_ART_NOT_EXIST = 2001

	// code = 3000... , category mode error
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "user exist",
	ERROR_PASSWORD_WRONG:     "password wrong",
	ERROR_USER_NOT_EXIST:     "user not exist",
	ERROR_TOKEN_NOT_EXIST:    "token not exist",
	ERROR_TOKEN_EXPIRED:      "token expired",
	ERROR_TOKEN_WRONG:        "token wrong",
	ERROR_TOKEN_WRONG_FORMAT: "token wrong format",
	ERROR_CATENAME_USED:      "category already exist",
	ERROR_CATE_NOT_EXIST:     "category not exist",
	ERROR_ART_NOT_EXIST:      "article not exist",
	ERROR_USER_NO_RIGHT:      "user doesn't have permission",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
