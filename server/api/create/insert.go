package create

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/database"
)

type Person struct {
	Username   string      `json:"username" bson:"username" binding:"required,min=1,max=16"`
	Nickname   string      `json:"nickname" bson:"nickname" binding:"required,min=1,max=16,nefield"`
	Age        json.Number `json:"age" bson:"age" binding:"required,min=1,max=100"`
	Password   string      `json:"password" bson:"password" binding:"required,min=1,max=16"`
	RePassword string      `json:"re_password" bson:"re_password" binding:"required,min=1,max=16,eqfield=Password"`
	Tel        string      `json:"tel" bson:"tel" binding:"required,len=11"`
}

func InsertOne(c *gin.Context) {
	// 处理数据库连接出现的错误
	if database.MongodbErr != nil {
		panic(database.MongodbErr)
	}

	// 定义用户结构体
	var person Person

	// 将前端传入的数据赋值给结构体, 如果校验失败, 则返回失败信息
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err,
		})
		panic(err)
		return
	}

	// 插入结构体数据
	result, insertErr := database.UsersColl.InsertOne(context.TODO(), person)
	if insertErr != nil {
		fmt.Println("insertErr", insertErr)
		panic(insertErr)
	}

	// 返回插入成功的消息与数据库的ObjectID, 可根据ObjectID查询到相应插入的数据
	c.JSON(http.StatusCreated, gin.H{
		"msg":      "OK",
		"ObjectID": result.InsertedID,
	})
}
