package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		//查看源码即可知道Query和DefaultQuery都是GetQuery
		//GetQuery方法的底层实现其实是c.Request.URL.Query().Get(key)，通过url.URL.Query()来获取所有的参数键值对。
		c.String(200, c.Query("wechat"))
		c.String(200, c.DefaultQuery("id", "0"))

		//类似的，我们也可以用QueryMap，QueryArray等查Map和数组相关数据
		c.JSON(200, c.QueryMap("ids"))
		c.JSON(200, c.QueryArray("media"))
	})

	r.POST("/", func(context *gin.Context) {
		//使用PostForm方法来获取相应的键值对，它接收一个key，也就是我们html中input这类表单标签的name属性值。
		wechat := context.PostForm("wechat")
		context.String(200, wechat)
	})

	//稍微总结下，gin框架在路由参数这块的查询上，大量用上了缓存getQueryCache来提高性能
	_ = r.Run(":8081")
}
