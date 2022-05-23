package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type registerForm struct{
	email	string `json:"email"`
	password string `json:"password"`
}

func register(c *gin.Context){
	//var json registerForm
	json := make(map[string]interface{})
	c.ShouldBindJSON(&json)
	fmt.Println(json["email"],json["password"])
}

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.POST("/register",register)
	return r
}