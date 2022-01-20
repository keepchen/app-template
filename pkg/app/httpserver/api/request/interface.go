package request

import "github.com/keepchen/app-template/pkg/constants"

type IRequest interface {
	//Validator 校验请求参数
	Validator() (constants.CodeType, error)
}
