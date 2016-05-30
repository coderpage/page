package status

const (
	UkownError          = -1    // 未知错误
	OK                  = 200   // OK
	ServerInternalErr   = 500   // 服务器内部错误
	UnprocessableEntity = 10001 // 无法处理的请求实体
	InvalidUserName     = 10002 // 无效用户名
	InvalidPwd          = 10003 // 无效密码
	UserExist           = 10004 // 该用户已存在
	UserNotExist        = 10005 // 用户不存在
	UserNotActivated    = 10006 // 用户未激活
	WrongUserNameOrPwd  = 10007 // 用户名或密码错误
	SendEmailFailed     = 10008 // 发送邮件失败
	TokenExpired        = 10009 // 令牌过期
	SourceNotFound      = 10010 // 未找到数据
	IllegalReqParam     = 10011 // 请求参数错误
)

var statusText = map[int]string{
	OK:                  "ok",
	ServerInternalErr:   "server internal error",
	UkownError:          "unkown error",
	UnprocessableEntity: "parse request entity failed",
	InvalidUserName:     "invalid user name",
	InvalidPwd:          "invalid password",
	UserExist:           "user already exist",
	UserNotExist:        "user not exist",
	UserNotActivated:    "user not activated",
	WrongUserNameOrPwd:  "user name or password is wrong",
	TokenExpired:        "token was expired",
	SourceNotFound:      "source not found on server",
}

func Text(code int) string {
	text := statusText[code]
	if text == "" {
		return "undefined error"
	}
	return text
}
