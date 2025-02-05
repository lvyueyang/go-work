# go 服务端

## 项目初次启动

### 添加配置文件

在 `server` 目录下添加 `config.dev.toml` 配置文件, 配置模板参考 `server/config/config.template.toml`

### 启动时初始化数据库表

```sh
go run main.go --init-db-model 1
```

> 仅首次启动或者数据库表有变化时添加 `--init-db-model` 即可

## swagger 文档生成

```sh
swag init --parseDependency
```

## 生成 crud

```sh
go run cmd/gen_module/gen.go --name work --cname 工作
```

## 添加数据库模型

db/db.go
Models 添加 model.Work

## 生成 dao

```sh
go run cmd/gen_dao/dao.go
```

## 初始化超级管理员用户

### 方式一

1. 启动 client-cms 前端
2. 打开 init-root-user 页面
3. 输入表单完成超级管理员初始化

### 方式二

参照 swagger 文档，调用 `/api/admin/auth/init-root-user` 接口完成初始化
