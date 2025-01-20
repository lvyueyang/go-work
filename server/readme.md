# go 服务端

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
