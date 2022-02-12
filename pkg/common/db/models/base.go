package models

import "time"

//BaseModel 基础模型字段
type BaseModel struct {
	ID        uint64     `gorm:"column:id;primary_key; AUTO_INCREMENT"`
	CreatedAt time.Time  `gorm:"column:created_at"` //创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at"` //更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at"` //(软)删除时间
}
