package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type reqParam struct {
	Name string  `json:"name"`
	Mail string  `json:"mail"`
}

func main() {
	router := gin.Default()

	router.POST("/", funcPost)

	// サーバーを起動
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal("サーバー起動に失敗", err)
	}
}

func funcPost (c *gin.Context) {
	//リクエストヘッダー
	ct := c.Request.Header.Get("Content-Type")

	//リクエストボディ
	var param reqParam
	c.BindJSON(&param)

	//クエリ
	id := c.Query("id")

	//レスポンス
	c.JSON(http.StatusOK, gin.H{
		"content-type": ct,
		"id": id,
		"name": param.Name,
		"mail": param.Mail,
	})
}
