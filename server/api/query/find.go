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

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func Find(c *gin.Context) {
	var results []User // 定义切片结构体(Find返回多个可能结果)

	key := c.Query("key")     // 定义查询的字段名
	value := c.Query("value") // 定义查询的字段值

	fmt.Println("key:", key)
	fmt.Println("value:", value)

	filter := bson.D{{key, value}} // 定义查询条件

	result, err := database.UsersColl.Find(context.TODO(), filter) // 查询

	// 处理错误, 如果没有错误则将result查询的结果赋值给结构体results
	if err = result.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	// 处理数据库连接出现的错误
	if database.MongodbErr != nil {
		panic(database.MongodbErr)
	}

	// 循环获取输出获取到的值
	for _, result1 := range results {
		res, _ := json.Marshal(result1)
		fmt.Println("res:", string(res))
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"result": results,
	})
}
