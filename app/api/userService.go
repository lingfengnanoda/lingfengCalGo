package api

import (
	"com.huahuo/app/dao/dal"
	"com.huahuo/app/dao/model"
	"com.huahuo/app/middle"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"time"
)

func login(ctx *gin.Context) {
	user := model.User{}
	ctx.ShouldBind(&user)
	u := dal.User
	query := u.Where(u.Name.Eq(user.Name))
	realUser, err := query.First()
	var id = realUser.ID
	claims := middle.CustomClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: math.MaxInt32, //永不过期
			Issuer:    "huahuo",
		},
	}
	j := middle.NewJWT()
	_, err = j.CreateToken(claims)
	if err != nil {
		log.Fatalln(err)
	}
	ctx.JSON(200, realUser)
}
func saveUser(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	log.Print(id) //设置接受的结构体
	u := dal.User
	//条件一
	query := u.Where(u.Name.Eq("花火"))
	//条件二
	query.Where(u.Age.Eq(21))
	//查询
	user, err := query.First()
	if err != nil {
		fmt.Print(err)
	}
	//返回
	ctx.JSON(200, user)
}
