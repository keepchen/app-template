package service

import (
	"github.com/keepchen/app-template/pkg/common/db/service"
	"github.com/keepchen/app-template/pkg/lib/db"
	"go.uber.org/zap"
)

//GetUserNameByID 根据用户编号获取用户名称
func GetUserNameByID(logger *zap.Logger, userID uint64) string {
	user, err := service.NewUserSvcImpl(logger, db.GetInstance()).GetUserByID(userID)
	if err != nil {
		return ""
	}

	return user.Username
}
