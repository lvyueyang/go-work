package types

// Pagination 分页查询
type Pagination struct {
	Current  int `json:"current" form:"current" binding:"omitempty,min=1" label:"当前页"`
	PageSize int `json:"page_size" form:"page_size" binding:"omitempty,min=0" label:"每页数量"`
}

// Order 排序
type Order struct {
	OrderKey  string `json:"order_key" form:"order_key" label:"被排序字段"`
	OrderType string `json:"order_type" form:"order_type" enums:"ase,desc" label:"排序方式" `
}
