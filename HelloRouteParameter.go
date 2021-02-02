package main

import "github.com/gin-gonic/gin"

func HelloRouteParameter() {
	r := gin.Default()

	///users/:id 就是一种路由匹配模式，也是一个通配符，其中:id就是一个路由参数，我们可以通过c.Param("id")获取定义的路由参数的值，然后用来做事情，比如打印出来。
	r.GET("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		context.String(200, "The user id is %s", id)
	})

	_ = r.Run(":8081")
}
