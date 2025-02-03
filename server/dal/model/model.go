package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey" validate:"required"` // ID
	CreatedAt time.Time      `json:"created_at" validate:"required"`           // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`                               // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`                           // 删除时间
}
