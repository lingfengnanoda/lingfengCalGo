package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	//user组
	user := r.Group("/user")
	user.POST("/save", saveUser)
}
