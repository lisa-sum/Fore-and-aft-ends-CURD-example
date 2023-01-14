package mutations

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
	Update string `json:"update" bson:"update"`
}

func UpdateOne(c *gin.Context) {
	var verify Verify
	if err := c.ShouldBind(&verify); err != nil {
		fmt.Println("转换失败")
		panic(err)
	}

	filter := bson.D{{verify.Key, verify.Value}}
	update := bson.D{{"$set", bson.D{{verify.Key, verify.Update}}}}

	// 对输入的查询条件进行检索,如果文档(表)有对应的结果, 则更新
	result, err := database.UsersColl.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "更新失败",
		})
		panic(err)
		return
	}

	// 更新的文档数
	fmt.Printf("Number of documents updated: %v\n", result.ModifiedCount)

	// 如果与数据库文档的字段或字段值不匹配,则返回自定义消息
	if result.ModifiedCount == 0 {
		fmt.Println("请检查输入的key与value")
		c.JSON(http.StatusResetContent, gin.H{
			"msg": "数据库文档没有该字段或者字段值, 请检查输入的key与value", // 返回更新的状态
		})
		return
	}

	// 更新成功返回的消息
	c.JSON(http.StatusOK, gin.H{
		"msg":     "更新成功",               // 返回更新的状态
		"updated": result.ModifiedCount, // 更新的文档数
	})
}
