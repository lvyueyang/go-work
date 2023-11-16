package permission

const (
	Admin{{.Name}}Find       string = "admin:{{.DbName}}:find:list"
	Admin{{.Name}}FindDetail string = "admin:{{.DbName}}:find:detail"
	Admin{{.Name}}Create     string = "admin:{{.DbName}}:create"
	Admin{{.Name}}UpdateInfo string = "admin:{{.DbName}}:update:info"
	Admin{{.Name}}Delete     string = "admin:{{.DbName}}:delete"
)

var AdminLabelMap = map[string]LabelType{
	Admin{{.Name}}Find: {
		Label: "查询{{.CName}}列表",
	},
	Admin{{.Name}}FindDetail: {
		Label: "查询{{.CName}}详情",
	},
	Admin{{.Name}}Create: {
		Label: "创建{{.CName}}",
	},
	Admin{{.Name}}UpdateInfo: {
		Label: "修改{{.CName}}信息",
	},
	Admin{{.Name}}Delete: {
		Label: "删除{{.CName}}",
	},
}
