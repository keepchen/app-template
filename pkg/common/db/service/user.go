package service

import (
	"github.com/keepchen/app-template/pkg/common/db/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//UserSvc 用户service接口
type UserSvc interface {
	GetUserList(page, count int) ([]*models.User, error)
	GetUserByID(userID uint64) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

//UserSvcImpl 用户service实现
type UserSvcImpl struct {
	logger     *zap.Logger
	dbInstance *gorm.DB
}

//NewUserSvcImpl 获取用户service实现实例
var NewUserSvcImpl = func(logger *zap.Logger, dbInstance *gorm.DB) *UserSvcImpl {
	return &UserSvcImpl{
		logger:     logger,
		dbInstance: dbInstance,
	}
}

//GetUserList 获取用户列表
func (u *UserSvcImpl) GetUserList(pageNumber, pageSize int) ([]*models.User, error) {
	var userList []*models.User
	if pageNumber < 1 {
		pageNumber = 0
	} else {
		pageNumber -= 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 100
	}
	err := u.dbInstance.Model(&models.User{}).Limit(pageSize).Offset(pageNumber * pageSize).Find(&userList).Error

	return userList, err
}

//GetUserByID 通过id获取用户信息
func (u *UserSvcImpl) GetUserByID(userID uint64) (*models.User, error) {
	var user *models.User
	err := u.dbInstance.Model(&models.User{}).Where(&models.User{
		BaseModel: models.BaseModel{
			ID: userID,
		},
	}).First(user).Error

	return user, err
}

//GetUserByUsername 通过username获取用户信息
func (u *UserSvcImpl) GetUserByUsername(username string) (*models.User, error) {
	var user *models.User
	err := u.dbInstance.Model(&models.User{}).Where(&models.User{
		Username: username,
	}).First(user).Error

	return user, err
}
