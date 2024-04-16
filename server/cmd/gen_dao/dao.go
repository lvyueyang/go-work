package main

import (
	"server/config"
	"server/db"

	"gorm.io/gen"
)

func main() {
	config.New()
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal/dao",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	database := db.Connect(true)

	g.UseDB(database)

	g.ApplyBasic(db.Models...)
	g.ApplyInterface(func(CommonFindQuery) {}, db.Models...)

	g.Execute()
}

type CommonFindQuery interface {

	// FindByID // 根据 ID 查询
	//
	// SELECT * FROM @@table WHERE id=@id
	FindByID(id uint) (*gen.T, error)
}
