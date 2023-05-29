package api

import (
	"com.huahuo/app/dao"
	"com.huahuo/app/dao/dal"
	"fmt"
	"github.com/gin-gonic/gin"
)

func saveUser(ctx *gin.Context) {
	//设置连接的DB
	dal.SetDefault(dao.DB)

	//设置接受的结构体
	u := dal.User
	//条件一
	query := u.Where(u.Name.Eq("花火"))
	//条件二
	query.Where(u.Age.Eq(21))
	//查询
	user, err := query.Find()
	if err != nil {
		fmt.Print(err)
	}
	//返回
	ctx.JSON(200, user)
}
