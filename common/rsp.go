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
	// CLIENT 客户端错误一级宏观错误码
	CLIENT Code = iota + 10000
)

const (
	// ParameterError 用户请求参数错误 二级宏观错误码
	ParameterError Code = iota + 10100
	// RequiredParameterIsEmpty 请求必填参数为空
	RequiredParameterIsEmpty
	// InvalidParameterValue 参数值非法
	InvalidParameterValue
)

const (
	// UserError 用户错误二级宏观错误码
	UserError Code = iota + 10200
	// UserNotRegister 用户未注册
	UserNotRegister
)

const (
	// CloudError 云环境错误二级宏观错误码
	CloudError Code = iota + 10300
	// OpenDataFail 获取OpenData失败
	OpenDataFail
)

// 定义errorCode对应的文本信息
var errorMsg = map[Code]string{
	SUCCESS:                  "成功",
	CLIENT:                   "用户端错误",
	UserError:                "用户错误",
	UserNotRegister:          "用户未注册",
	CloudError:               "云环境错误",
	OpenDataFail:             "获取OpenData失败",
	ParameterError:           "用户请求参数错误",
	RequiredParameterIsEmpty: "请求必填参数为空",
}

// CodeText 根据错误码获取错误信息
func CodeText(code Code) string {
	return errorMsg[code]
}

// Rsp 响应
type Rsp struct {
	Code  Code        `json:"code"`            // 错误码
	Msg   string      `json:"msg,omitempty"`   // 消息
	Data  interface{} `json:"data,omitempty"`  // 数据
	Total int64       `json:"total,omitempty"` // 数据
}

func (e *Rsp) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, advice: %s", e.Code, e.Msg, e.Msg)
}

// Ok 成功的消息返回
func Ok() Rsp {
	return Rsp{Code: SUCCESS, Msg: CodeText(SUCCESS)}
}

// OkData 成功的数据返回
func OkData(data interface{}) Rsp {
	return Rsp{Code: SUCCESS, Msg: CodeText(SUCCESS), Data: data}
}

// OkPage 成功的数据返回
func OkPage(total int64, data interface{}) Rsp {
	return Rsp{Code: SUCCESS, Msg: CodeText(SUCCESS), Data: data, Total: total}
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
