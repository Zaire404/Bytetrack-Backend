package response

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	//成员错误
	ErrorExistUser      = 10002
	ErrorNotExistUser   = 10003
	ErrorFailEncryption = 10006
	ErrorNotCompare     = 10007

	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorDatabase              = 40001
)

var MsgFlags = map[int]string{
	SUCCESS:       "操作成功",
	ERROR:         "操作失败",
	InvalidParams: "请求参数错误",

	ErrorExistUser:    "用户已存在",
	ErrorNotExistUser: "用户不存在",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorNotCompare:            "不匹配",
	ErrorDatabase:              "数据库操作出错,请重试",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// RespSuccess 成功返回
func RespSuccess(data string) *Response {
	status := SUCCESS
	r := &Response{
		Status: status,
		Data:   data,
		Msg:    GetMsg(status),
	}

	return r
}

// RespSuccessWithData 带data成功返回
func RespSuccessWithData(data interface{}) *Response {
	status := SUCCESS

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    GetMsg(status),
	}

	return r
}

// RespError 错误返回
func RespError(err error, data string) *Response {
	status := ERROR

	r := &Response{
		Status: status,
		Msg:    GetMsg(status),
		Data:   data,
		Error:  err.Error(),
	}

	return r
}
