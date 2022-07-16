package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirwarj/go-api-gin/handle"
	"github.com/sirwarj/go-api-gin/models"
)

func AllData(m *handle.MemberData) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(200, m.AllData())
	}
}

func AddData(m *handle.MemberData) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var v models.Member

		if err := ctx.ShouldBindJSON(&v); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		m.AddData(v)
		ctx.JSON(200, v)
	}
}
