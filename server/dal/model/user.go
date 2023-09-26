package model

import (
	"server/consts"
)

type User struct {
	BaseModel
	Name      string            `json:"name"`
	Age       uint32            `json:"age"`
	Email     string            `json:"email" gorm:"unique"`
	Avatar    string            `json:"avatar"`
	Status    consts.UserStatus `json:"status" gorm:"default=1"`
	Username  string            `json:"username" gorm:"unique"`
	Password  string            `json:"-"`
	WxOpenId  string            `json:"wx_open_id"`
	WxUnionId string            `json:"wx_union_id"`
}

func (User) TableName() string {
	return "user"
}
