package db

import (
	"fmt"
	"reflect"
	"server/config"
	"server/dal/dao"
	"server/dal/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

// Models 数据库表
var Models = []any{
	model.User{},
	model.Captcha{},
	model.AdminUser{},
	model.AdminRole{},
	// 附加业务
	model.News{},
}

func New(initModel bool) {
	Database = Connect(initModel)
}

// Connect 数据库连接
func Connect(initModel bool) *gorm.DB {
	fmt.Println("数据库连接中")
	now := time.Now()
	dsn := config.Config.Db.Link

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}

	dao.SetDefault(db)

	if initModel {
		for _, m := range Models {
			structType := reflect.TypeOf(m)

			if err := db.AutoMigrate(m); err != nil {
				fmt.Printf("表 %+v 初始化失败\n", structType.Name())
				panic(err)
			} else {
				fmt.Printf("表 %+v 初始化成功\n", structType.Name())
			}
		}

	}
	fmt.Printf("数据库连接成功，耗时 %+v\n", time.Now().Sub(now))
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	fmt.Printf("\n%+v\n", sqlDB.Stats())
	return db
}

func Close() {
	sqlDB, _ := Database.DB()
	fmt.Printf("数据库关闭中\n%+v\n", sqlDB.Stats())
	sqlDB.Close()
	fmt.Printf("数据库连接已关闭\n%+v\n", sqlDB.Stats())
}
