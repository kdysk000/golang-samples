package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gorm-sample/model"
	"gorm-sample/mysql"

	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	b, _ := ioutil.ReadAll(c.Request.Body)
	var user model.User
	json.Unmarshal(b, &user)
	err := mysql.AddUser(user)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func UserList(c *gin.Context) {
	users, err := mysql.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, users)
}

func UserUpdate(c *gin.Context) {
	b, _ := ioutil.ReadAll(c.Request.Body)
	var user model.User
	json.Unmarshal(b, &user)
	err := mysql.UpdateUser(user)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func UserDelete(c *gin.Context) {
	userid := c.Query("id")
	err := mysql.DeleteUser(userid)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
