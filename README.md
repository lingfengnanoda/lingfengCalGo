# HuaHuoGiao

> 基于gin，gorm gen的轻量级易用goweb项目整合脚手架，在go语言下体验springboot+mybatis-plus的懒人开发

## 项目结构

![image-20230530211348025](http://image.fzuhuahuo.cn/image-20230530211348025.png)

## 一.init项目

### 1.下载gin

```shell
go get -u github.com/gin-gonic/gin
```

### 2.下载gorm mysql驱动 gorm gen

```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u gorm.io/gen
```

### 3.配置mysql连接

> dao包下面的gorm.go，最上方的连接参数字符串改一下就可以了

### 4.gorm gen 生成代码

>运行gen下的main.go

## 二.CRUD使用方法

#### 条件拼接查询

```go
func saveUser(ctx *gin.Context) {
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

```

#### 分页查询

```go
userList, count, err := u.Where(u.ID.Gt(2)).Order(u.Age.Desc()).FindByPage(1, 1)
```

## 三.中间件

- jwt权限认证

### 
