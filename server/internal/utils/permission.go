package utils

import "fmt"

type PermissionInfo struct {
	// 权限码
	Code string `json:"code" binding:"required"`
	// 权限名称
	Name string `json:"name" binding:"required"`
}

var PermissionList = []PermissionInfo{}

func CreatePermission(code string, name string) PermissionInfo {
	// 判断 code 在 PermissionList 中是否存在
	for _, permission := range PermissionList {
		if permission.Code == code {
			// 终止程序
			message := fmt.Sprintf("权限码 %s 已存在", code)
			panic(message)
		}
	}
	for _, permission := range PermissionList {
		if permission.Name == name {
			// 终止程序
			message := fmt.Sprintf("权限名称 %s 已存在", code)
			panic(message)
		}
	}

	info := PermissionInfo{
		Code: code,
		Name: name,
	}

	PermissionList = append(PermissionList, info)

	return info
}

func FindPermission(code string) *PermissionInfo {
	for _, permission := range PermissionList {
		if permission.Code == code {
			return &permission
		}
	}
	return &PermissionInfo{}
}
