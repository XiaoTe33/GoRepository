package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/info1", func(context *gin.Context) {
		context.JSON(200, "get a info ")
	})
	r.POST("/info2", func(context *gin.Context) {
		context.JSON(200, "create a info ")
	})
	r.DELETE("/info3", func(context *gin.Context) {
		context.JSON(200, "delete a info ")
	})
	r.PUT("/info4", func(context *gin.Context) {
		context.JSON(200, "update a info ")
	})
	r.Run(":8000")
}
