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
	// code = 2000... , article mode error

	// code = 3000... , category mode error

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
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
