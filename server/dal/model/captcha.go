package model

import (
	"server/internal/consts"
	"time"
)

type Captcha struct {
	ID          uint                 `json:"id" gorm:"primarykey" validate:"required"`
	CreatedAt   time.Time            `json:"created_at" validate:"required"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Expiration  time.Time            `json:"expiration" validate:"required"` // 过期时间
	Current     string               `json:"current" validate:"required"`    // 所属目标 手机号/邮箱
	CurrentType consts.CaptchaType   `json:"current_type" validate:"required"`
	Code        string               `json:"code" validate:"required"`
	Status      consts.CaptchaStatus `json:"status" validate:"required"`
	Scenes      consts.CaptchaScenes `json:"scenes" validate:"required"` // 使用场景
}

func (Captcha) TableName() string {
	return "captcha"
}
