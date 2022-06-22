package response

import (
	"encoding/xml"
	"fmt"

	"github.com/keepchen/app-template/pkg/constants"
)

type IXMLResponse interface {
	//Assemble 组装返回值给客户端
	Assemble(code constants.CodeType, data interface{}) XMLResponse
	//GetData 获取数据
	GetData() string
}

//XMLResponse XML返回结构
//
/*
<response>
	<code>0</code>
	<msg></msg>
	<success>true</success>
	<data><![CDATA[...]]></data>
</response>
*/
type XMLResponse struct {
	XMLName xml.Name `xml:"response"`
	Code    int      `xml:"code"`       //错误码
	Msg     string   `xml:"msg"`        //错误信息
	Success bool     `xml:"success"`    //是否成功
	Data    string   `xml:"data,cdata"` //返回数据
}

//Assemble 组装返回值
func (xr *XMLResponse) Assemble(code constants.CodeType, data interface{}) XMLResponse {
	xr.Code = code.Int()
	xr.Msg = code.String()
	if code == constants.ErrNone {
		xr.Success = constants.Success
	} else {
		xr.Success = constants.Fail
	}
	xr.Data = fmt.Sprintf("%v", data)

	return *xr
}

//GetData 获取数据
func (xr *XMLResponse) GetData() string {
	if output, err := xml.MarshalIndent(xr, "", "    "); err == nil {
		return xml.Header + string(output)
	}

	return ""
}
