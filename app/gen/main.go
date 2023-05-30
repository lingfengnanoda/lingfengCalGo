package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

func main() {
	user := "root"
	password := "20030416cjh"
	hostAndPort := "tycloud.fzuhuahuo.cn:3308"
	database := "gotest"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, hostAndPort, database)
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
