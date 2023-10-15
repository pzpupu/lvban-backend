package common

import "fmt"

// Code 响应代码
type Code int

// 定义响应代码常量。
const (
	// SUCCESS 成功
	SUCCESS Code = iota
)

const (
	// CLIENT 一级宏观错误码
	CLIENT Code = iota + 10001
	// UserNotRegister 用户未注册
	UserNotRegister
)

// 定义errorCode对应的文本信息
var errorMsg = map[Code]string{
	SUCCESS:         "成功",
	CLIENT:          "用户端错误",
	UserNotRegister: "用户未注册",
}

// CodeText 根据错误码获取错误信息
func CodeText(code Code) string {
	return errorMsg[code]
}

// Rsp 响应
type Rsp struct {
	Code Code        `json:"code,omitempty"` // 错误码
	Msg  string      `json:"msg,omitempty"`  // 消息
	Data interface{} `json:"data,omitempty"` // 数据
}

func (e *Rsp) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, advice: %s", e.Code, e.Msg, e.Msg)
}

// OkData 成功的数据返回
func OkData(data interface{}) Rsp {
	return Rsp{Code: SUCCESS, Msg: CodeText(SUCCESS), Data: data}
}

// OkMsg 成功的消息返回
func OkMsg(msg string) Rsp {
	return Rsp{Code: SUCCESS, Msg: msg}
}

// FailMsg 失败的消息返回
func FailMsg(code Code, msg string) Rsp {
	return Rsp{Code: code, Msg: msg}
}

// Fail 失败的消息返回
func Fail(code Code) Rsp {
	return Rsp{Code: code, Msg: CodeText(code)}
}
