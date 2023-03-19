package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postBody struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

type getUrl struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/", funcPost)
	router.GET("/:name/:id", funcGet)

	// サーバーを起動
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal("サーバー起動に失敗", err)
	}
}

// ヘッダー、ボディ、クエリ
func funcPost(c *gin.Context) {
	//リクエストヘッダー
	ct := c.Request.Header.Get("Content-Type")

	//リクエストボディ
	b, _ := ioutil.ReadAll(c.Request.Body)
	var param postBody
	json.Unmarshal(b, &param)

	//クエリ
	id := c.Query("id")

	//レスポンス
	c.JSON(http.StatusOK, gin.H{
		"content-type": ct,
		"id":           id,
		"name":         param.Name,
		"mail":         param.Mail,
	})
}

// URLバインド
func funcGet(c *gin.Context) {
	var param getUrl
	c.ShouldBindUri(&param)
	c.JSON(http.StatusOK, gin.H{
		"id":   param.ID,
		"name": param.Name,
	})

}
