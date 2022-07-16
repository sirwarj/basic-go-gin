package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirwarj/go-api-gin/controllers"
	"github.com/sirwarj/go-api-gin/handle"
)

func main() {
	r := gin.Default()
	r.Use(Authorization)

	m := handle.NewMember()

	r.GET("/member", controllers.AllData(m))

	r.POST("/member", controllers.AddData(m))

	r.Run()
}

func Authorization(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	auth := strings.Split(token, " ")

	if len(auth) != 2 || auth[0] != "Bearer" {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Authorization failed"})
		return
	}

	if auth[1] != "asdoaijsdohasad" {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid Token"})
		return
	}

	ctx.Next()
}
