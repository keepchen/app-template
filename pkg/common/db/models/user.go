package models

//User 用户表
type User struct {
	BaseModel
	Username string `gorm:"column:username"` //用户名
	Password string `gorm:"column:password"` //密码
}

//TableName 表名称
func (*User) TableName() string {
	return "data_users"
}
