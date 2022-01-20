package response

import "github.com/keepchen/app-template/pkg/constants"

type IResponse interface {
	//Assemble 组装返回值给客户端
	Assemble(code constants.CodeType, data interface{}) StandardResponse
}

//StandardResponse 标准返回结构
type StandardResponse struct {
	Code    constants.CodeType `json:"code"`    //错误码
	Msg     string             `json:"msg"`     //错误信息
	Success bool               `json:"success"` //是否成功
	Data    interface{}        `json:"data"`    //返回数据
}

//Assemble 组装返回值
func (sr *StandardResponse) Assemble(code constants.CodeType, data interface{}) StandardResponse {
	sr.Code = code
	sr.Msg = code.String()
	if code == constants.ErrNone {
		sr.Success = constants.Success
	} else {
		sr.Success = constants.Fail
	}
	sr.Data = data

	return *sr
}
