package api

import (
	"com.huahuo/app/middle"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/login", login)
	//user组
	user := r.Group("/user")
	user.Use(middle.JWTAuth())
	user.POST("/save", saveUser)

}
