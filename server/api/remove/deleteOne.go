package remove

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"server/database"
)

type Verify struct {
	Key    string `json:"key" bson:"key"`
	Value  string `json:"value" bson:"value"`
	Delete string `json:"delete" bson:"delete"`
}

func DeleteOne(c *gin.Context) {
	// 定义结构体
	var verify Verify

	// 绑定结构体
	err := c.ShouldBind(&verify)
	if err != nil {
		panic(err)
	}

	// 删除条件
	filter := bson.D{{verify.Key, verify.Value}}
	// 如果删除异常则抛出, 否则为成功的结果
	result, err2 := database.UsersColl.DeleteOne(context.TODO(), filter)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "失败",
		})
		panic(err2)
		return
	}

	// 打印删除条目, 如果filter的查询条件为0,则数据库没有该条目
	fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"msg": "要删除的条目在数据库不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    "成功",
		"result": result.DeletedCount,
	})
}
