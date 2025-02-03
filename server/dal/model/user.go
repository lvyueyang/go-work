package model

import (
	"server/internal/consts"
)

type User struct {
	BaseModel
	Name      string            `json:"name"`
	Age       uint32            `json:"age"`
	Email     string            `json:"email" gorm:"unique" validate:"required"`
	Avatar    string            `json:"avatar"`
	Status    consts.UserStatus `json:"status" gorm:"default=1" validate:"required"`
	Username  string            `json:"username" gorm:"unique" validate:"required"`
	Password  string            `json:"-"`
	WxOpenId  string            `json:"wx_open_id"`
	WxUnionId string            `json:"wx_union_id"`
}

func (User) TableName() string {
	return "user"
}
