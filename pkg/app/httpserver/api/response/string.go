package response

import (
	"encoding/json"

	"github.com/keepchen/app-template/pkg/constants"
)

type IStringResponse interface {
	//Assemble 组装返回值给客户端
	Assemble(code constants.CodeType, data interface{}) StringResponse
	//GetData 获取数据
	GetData() string
}

//StringResponse 字符串返回结构
type StringResponse struct {
	RawString string
}

func (sr *StringResponse) Assemble(_ constants.CodeType, data interface{}) StringResponse {
	if str, ok := data.(string); ok {
		sr.RawString = str
	} else {
		if byt, err := json.Marshal(data); err == nil {
			sr.RawString = string(byt)
		}
	}

	return *sr
}

//GetData 获取数据
func (sr *StringResponse) GetData() string {
	return sr.RawString
}
