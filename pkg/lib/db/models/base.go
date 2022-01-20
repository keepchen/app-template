package models

//BaseModel 基础模型字段
type BaseModel struct {
	ID        uint64 `gorm:"column:id;primary_key; AUTO_INCREMENT"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}
