package error

const CodeKey = "code"

const (
	OK              = 0
	BadRequest      = 400
	ServerError     = 500
	Unauth          = 4001
	StoreImageFail  = 5001
	ImageNotFound   = 5002
	FetchImageFail  = 5003
	ImageResp       = 5004
	ValidateSession = 5005
	SendCode        = 5006
	InvalidCode     = 4002
	UsernameDup     = 4003
	LoginError      = 4004
	AccountExists   = 4005
	PingUnauth      = 4006
	SelfOrder = 4007
	BookOrdered = 4008
)

var msg = map[int]string{
	OK:              "ok",
	BadRequest:      "请求参数不合法",
	StoreImageFail:  "上传图像失败",
	ImageNotFound:   "未能找到图像",
	FetchImageFail:  "获取图像失败",
	ImageResp:       "返回图像失败",
	Unauth:          "你还未登录",
	ValidateSession: "验证session失败",
	SendCode:        "发送验证码失败",
	InvalidCode:     "验证码非法",
	UsernameDup:     "用户名已经存在",
	LoginError:      "用户名或者密码错误",
	ServerError:     "服务异常",
	AccountExists:   "账户已存在",
	PingUnauth:      "Ping Unauth",
	SelfOrder: "不能购买/出售自己发布的书籍",
	BookOrdered: "书籍已经被人下单",
}

func Translate(code int) string {
	if val, ok := msg[code]; ok {
		return val
	}
	return "unknown code"
}
