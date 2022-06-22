package constants

import "fmt"

//ICodeType 错误码类型接口
type ICodeType interface {
	String() string
	Int() int
}

//CodeType 错误码类型
type CodeType int

//错误码
const (
	ErrNone                      = CodeType(0)            //没有错误，占位
	ErrRequestParamsInvalid      = CodeType(iota) + 10000 //请求参数有误
	ErrAuthorizationTokenInvalid                          //令牌已失效
)

//错误码信息表
//
//READONLY for concurrency safety
var codeMsgFill = map[CodeType]string{
	ErrNone:                      "",
	ErrRequestParamsInvalid:      "请求参数有误",
	ErrAuthorizationTokenInvalid: "令牌已失效",
}

//String 获取错误信息字符
func (ct CodeType) String() string {
	msg, ok := codeMsgFill[ct]
	if ok {
		return msg
	}

	return fmt.Sprintf("[Warn] ErrorCode {%d} not defined!", ct)
}

//Int 获取错误码
func (ct CodeType) Int() int {
	return int(ct)
}
