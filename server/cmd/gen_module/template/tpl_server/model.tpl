package model

import "time"

type {{.Name}} struct {
	BaseModel
    Cover     string    `json:"cover"`                        // 封面
    Title     string    `json:"title" gorm:"unique"`          // 标题
    Desc      string    `json:"desc"`                         // 描述
    Content   string    `json:"content" gorm:"type:longtext"` // 内容
    PushDate  time.Time `json:"push_date"`                    // 发布日期
    Recommend uint      `json:"recommend"`                    // 推荐等级
    AuthorID  uint      `json:"author_id"`                    // 作者
}

func (*{{.Name}}) TableName() string {
	return "{{.DbName}}"
}
