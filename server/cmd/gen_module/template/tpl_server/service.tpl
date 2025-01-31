package service

import (
	"server/dal/dao"
	"server/dal/model"
	"server/lib/errs"
	"server/internal/types"
	"server/internal/utils"
	"strconv"
	"strings"
)

type {{.Name}}Service struct {
}

var {{.ServiceName}}Service = new({{.Name}}Service)

func New{{.Name}}Service() *{{.Name}}Service {
	return {{.ServiceName}}Service
}

type Find{{.Name}}ListOption struct {
	types.Pagination
	types.Order
	Keyword string `json:"keyword" form:"keyword"`
}

func (s *{{.Name}}Service) FindList(query Find{{.Name}}ListOption) (utils.ListResult[[]*model.{{.Name}}], error) {
	result := utils.ListResult[[]*model.{{.Name}}]{}
	n := dao.{{.Name}}
	q := n.Omit(n.Content).Where(
		n.Title.Like("%" + query.Keyword + "%"),
	)
	if id, err := strconv.ParseUint(query.Keyword, 10, 64); err == nil {
		q = q.Or(n.ID.Eq(uint(id)))
	}

	if query.OrderKey != "" {
		col, _ := n.GetFieldByName(query.OrderKey)
		if strings.ToLower(query.OrderType) == "desc" {
			q = q.Order(col.Desc())
		} else {
			q = q.Order(col)
		}
	} else {
		q.Order(n.CreatedAt.Desc())
	}

	if list, total, err := q.FindByPage(utils.PageTrans(query.Pagination)); err != nil {
		return result, err
	} else {
		result.List = list
		result.Total = total
	}

	return result, nil
}

func (s *{{.Name}}Service) FindDetail(id uint) (*model.{{.Name}}, error) {
	current, err := dao.{{.Name}}.FindByID(id)
	if err != nil {
		return current, errs.CreateServerError("{{.CName}}不存在", err, nil)
	}
	return current, err
}

func (s *{{.Name}}Service) Create(info model.{{.Name}}) (*model.{{.Name}}, error) {
	oldData, err := dao.{{.Name}}.Where(dao.{{.Name}}.Title.Eq(info.Title)).Take()

	if err == nil {
		return oldData, errs.CreateServerError("{{.CName}}名称重复", err, nil)
	}

	data := &model.{{.Name}}{
		Cover:     info.Cover,
		Title:     info.Title,
		Desc:      info.Desc,
		Content:   info.Content,
		PushDate:  info.PushDate,
		Recommend: info.Recommend,
		AuthorID:  info.AuthorID,
	}
	if err := dao.{{.Name}}.Create(data); err != nil {
		return oldData, err
	}

	return data, nil
}

func (s *{{.Name}}Service) Update(id uint, info model.{{.Name}}) (*model.{{.Name}}, error) {
	current, err := dao.{{.Name}}.FindByID(id)
	if err != nil {
		return current, errs.CreateServerError("{{.CName}}不存在", err, nil)
	}
	data := &model.{{.Name}}{
		Cover:     info.Cover,
		Title:     info.Title,
		Desc:      info.Desc,
		Content:   info.Content,
		PushDate:  info.PushDate,
		Recommend: info.Recommend,
	}

	if _, err := dao.{{.Name}}.Where(dao.{{.Name}}.ID.Eq(id)).Updates(&data); err != nil {
		return current, err
	}
	return data, nil
}

func (s *{{.Name}}Service) Delete(id uint) error {
	if _, err := dao.{{.Name}}.FindByID(id); err != nil {
		return errs.CreateServerError("{{.CName}}不存在", err, nil)
	}

	if _, err := dao.{{.Name}}.Where(dao.{{.Name}}.ID.Eq(id)).Delete(); err != nil {
		return err
	}
	return nil
}
