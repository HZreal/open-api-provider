package common

type ResponseCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var ParamsError = ResponseCode{Code: 94000, Msg: "参数错误"}

var NotLogin = ResponseCode{94010, "未登录!"}

var NoAuth = ResponseCode{94011, "无权限!"}

var ServerError = ResponseCode{95000, "服务器错误!"}

var RedisError = ResponseCode{950001, "Redis错误!"}
var MysqlError = ResponseCode{950002, "Mysql错误!"}
