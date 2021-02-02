package main

import "github.com/gin-gonic/gin"

func HelloGin() {
	//实例化一个默认的gin实例
	r := gin.Default()

	//以GET访问注册的一个处理函数，相对地址为/
	r.GET("/", func(context *gin.Context) {
		//返回一个JSON格式的字符串
		//code则为请求成功的状态码，若是其他码也可以写
		//gin.H为一个map接口函数，在JSON函数中为的obj
		//gin.H--map[string]interface{}，声明为H类型，我们可以点开源码查看
		context.JSON(200, gin.H{
			"github": "https://github.com/6ylqq",
			"weChat": "10086",
		})
	})

	//启动一个HTTP服务，端口为8081
	//注意指定返回值，返回空
	_ = r.Run(":8081")
}
