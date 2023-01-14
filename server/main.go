package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/api/create"
	"server/api/mutations"
	"server/api/query"
	"server/api/remove"
	"server/database"
)

const PORT = ":4000" // 服务端口号, 如果与先有端口冲突可替换为任意端口尝试

func main() {
	server := gin.Default() // 启动Gin HTTP服务
	server.Use(cors())      // 处理跨域的中间件

	// 用户的路由组
	userGroup := server.Group("/user")
	{
		userGroup.GET("/find", query.Find)              // 查询路由
		userGroup.POST("/login", query.FindOne)         // 登录校验路由
		userGroup.PUT("/create", create.InsertOne)      // 增加路由
		userGroup.PATCH("/update", mutations.UpdateOne) // 更新路由
		userGroup.DELETE("/delete", remove.DeleteOne)   // 删除路由
	}

	err := server.Run(PORT)
	if err != nil {
		return
	}
	fmt.Println("服务启动成功!")

	// 如果连接数据库失败, 则抛出异常
	if database.MongodbErr != nil {
		panic(database.MongodbErr.Error())
	}
	fmt.Println("成功连接并ping通Mongodb服务!")
	// 使用完数据库关闭连接
	defer func() {
		if err := database.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

// 处理跨域
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE,PATCH")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
			//
			//c.Header("Content-Type", "application/json")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //自定义 Header
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(http.StatusNoContent)
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()
		c.Next()
	}
}
