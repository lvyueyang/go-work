package main

import (
	"flag"
	"fmt"
	"net/http"
	"server/app"
	"server/config"
	"server/consts"
	"server/db"
	_ "server/docs"
	"server/lib/logger"
	"server/lib/run_server"
	"server/lib/valid"
	"server/middleware"
	"server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @title		男生自用 API 接口文档
// @version	1.0
func main() {
	initDBModel := flag.Bool("init-db-model", false, "初始化数据库表")
	fmt.Println("Version:", consts.Version)
	now := time.Now()

	// 配置
	configPath := flag.String("c", "config/config.dev.toml", "配置文件路径")
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
