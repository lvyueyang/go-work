package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"server/app"
	"server/config"
	"server/db"
	_ "server/docs"
	"server/internal/consts"
	"server/internal/middleware"
	"server/internal/utils"
	"server/lib/logger"
	"server/lib/run_server"
	"server/lib/valid"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed assets/*
var fe embed.FS

// @title		男生自用 API 接口文档
// @version	1.0
func main() {
	initDBModel := flag.Bool("init-db-model", false, "初始化数据库表")
	configPath := flag.String("c", "config.dev.toml", "配置文件路径")
	flag.Parse()

	fmt.Println("Version:", consts.Version)
	now := time.Now()

	// 配置
	config.New(*configPath)

	var envName = utils.EnumLabel(consts.EnvDev)
	if config.Config.IsProd {
		envName = utils.EnumLabel(consts.EnvProd)
	}

	// 日志
	logger.New()

	if config.Config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	// 数据库
	db.New(*initDBModel)
	defer db.Close()

	// gin
	router := gin.New()
	// 全局中间件
	router.Use(middleware.RequestLogger(), gin.Recovery())

	// 前端模板
	templ := template.Must(template.New("").ParseFS(fe, "assets/fe/*.html"))
	router.SetHTMLTemplate(templ)
	router.StaticFS("/public", http.FS(fe))
	// 验证器
	valid.New()

	// 启动
	app.New(router)

	s := run_server.New(&http.Server{
		Addr:    ":" + strconv.Itoa(config.Config.Port),
		Handler: router,
	})

	fmt.Println("\nAPI: ", "http://127.0.0.1:"+strconv.Itoa(config.Config.Port))
	fmt.Println("Swagger: ", "http://127.0.0.1:"+strconv.Itoa(config.Config.Port)+"/swagger/index.html")
	//fmt.Printf("配置文件加载成功 %+v\n", config.Config)
	fmt.Println("当前环境: ", envName)
	fmt.Println("启动耗时: ", time.Now().Sub(now))

	s.Run()
}

// func FS() http.FileSystem {
// 	return FePublic{}
// }

// type FePublic struct {
// }

// func (c FePublic) Open(name string) (http.File, error) {
// 	print(name + "\n")
// 	f, err := fe.Open("assets/fe-ui/" + name)
// 	return file, err
// }
