package router

import (
	"com.huahuo/app/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api.RegisterRouter(r)
}
