package model

import "time"

type {{.Name}} struct {
	BaseModel
	Cover     string    `json:"cover"`                                   // 封面
	Title     string    `json:"title" gorm:"unique" validate:"required"` // 标题
	Desc      string    `json:"desc"`                                    // 简介
	Content   string    `json:"content" gorm:"type:longtext"`            // 内容
	PushDate  time.Time `json:"push_date"`                               // 发布日期
	Recommend uint      `json:"recommend" validate:"required"`           // 推荐等级
	IsVisible bool      `json:"is_visible" validate:"required"`          // 可见性
	AuthorID  uint      `json:"author_id" validate:"required"`           // 作者
}

func (*{{.Name}}) TableName() string {
	return "{{.DbName}}"
}
