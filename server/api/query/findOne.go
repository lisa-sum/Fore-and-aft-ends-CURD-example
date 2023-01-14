package query

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"server/database"
)

// Person 用户结构
type Person struct {
	Username string `json:"username" bson:"username" binding:"required,min=1,max=20"` // 如果需要对数据传递参数, 必须使用bson 这个tag
	Password string `json:"password" binding:"required,min=1,max=20"`
}

// Verify 校验
type Verify struct {
	Password string `json:"password"`
}

func FindOne(c *gin.Context) {
	var person Person             // 定义用户结构体
	var verify Verify             // 定义校验结构体
	err2 := c.ShouldBind(&person) // 将json绑定到结构体
	if err2 != nil {
		panic(err2)
	}

	// 设置查询条件
	filter := bson.D{{"username", person.Username}}
	// 返回结果, 如果成功则将返回的bson结果绑定到结构体
	result := database.UsersColl.FindOne(context.TODO(), filter).Decode(&verify)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名或者密码输入错误!",
		})
		panic(result)
		return
	}
	// 处理数据库连接出现的错误
	if database.MongodbErr != nil {
		panic(database.MongodbErr)
	}

	// 将bson二进制json转为json
	res, _ := json.Marshal(verify)
	// 以字符串形式输出转换结果
	fmt.Println(string(res))

	// 校验
	if person.Password != verify.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "密码错误!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username": person.Username,
		"msg":      "登录成功",
	})
}
