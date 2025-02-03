package model

import (
	"server/internal/consts"
)

type AdminUser struct {
	BaseModel
	Name     string                 `json:"name" validate:"required"`
	Username string                 `json:"username" gorm:"unique" validate:"required"`
	Password string                 `json:"-"`
	Email    string                 `json:"email" gorm:"unique"`
	Avatar   string                 `json:"avatar"`
	IsRoot   bool                   `json:"is_root"` // 是否是超级管理员
	Status   consts.AdminUserStatus `json:"status" validate:"required"`
	Roles    []*AdminRole           `json:"roles" gorm:"many2many:admin_user_roles;"`
	News     []*News                `json:"news" gorm:"foreignKey:AuthorID"`
}

func (AdminUser) TableName() string {
	return "admin_user"
}
