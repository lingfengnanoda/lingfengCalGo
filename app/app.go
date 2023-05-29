package main

import (
	"com.huahuo/app/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
