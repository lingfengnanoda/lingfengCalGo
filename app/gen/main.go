package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

const dsn = "root:20030416cjh@tcp(tycloud.fzuhuahuo.cn:3308)/gotest?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      ".././app/dao/dal",
		ModelPkgPath: ".././app/dao/model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
