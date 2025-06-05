# 数字岩石力学智能分析平台-前端工程

## 开发与构建


### 初始化和更新子模块

```
git submodule init
git submodule update --remote
```

### 拉取并切换子模块到指定分支

```
git submodule foreach 'branch=$(git config -f $toplevel/.gitmodules submodule.$name.branch); if [ -n "$branch" ]; then git checkout $branch; git pull origin $branch; fi'
```

### 安装依赖

```
pnpm i
```

### 启动（开发模式）

```
npm start
```

### 构建

```
npm run build
```

## 架构介绍

1. 业务模块  
   1-1. 全局配置  
   1-2. 登录认证 1-3. 页面模块（可独立构建）  
    1-3-1. 首页  
    1-3-2. 流程化建模  
    1-3-3. 井巷模式  
    1-3-4. ....
2. 主应用基座  
   2-1. 整合所有模块打包为一个完整 APP
3. 基础包  
   3-1. consts 公共常量包  
   3-2. request 公共网络请求包  
   3-3. types 公共类型  
   3-4. ui 基础 UI 组件  
   3-5. utils 公共工具函数  
   3-6. ThreeEditor threejs 基础继承类

## 开发文档

[如何创建模块](./docs/创建模块.md)  
[模块独立开发与部署](./docs/模块独立开发与部署.md)
