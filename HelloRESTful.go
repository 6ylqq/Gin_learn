package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func HelloRESTful() {
	users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}
	r := gin.Default()
	r.GET("users", func(context *gin.Context) {
		context.JSON(200, users)
	})
	r.POST("/users", func(context *gin.Context) {
		//创建一个用户
	})
	r.DELETE("/usrs/123", func(context *gin.Context) {
		//删除ID为123的用户
	})
	r.PUT("/usrs/123", func(context *gin.Context) {
		//更新ID为123的用户
	})

	r.PATCH("/usrs/123", func(context *gin.Context) {
		//更新ID为123用户的部分信息
	})

	_ = r.Run(":8081")
}
