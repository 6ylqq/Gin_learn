package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//分组主要是对付版本升级，模块切分等需求的时候，用起来就很爽

	//V1版本的API
	v1Group := r.Group("/v1")
	v1Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v1/users")
	})
	v1Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v1/products")
	})

	//分组的嵌套
	v1AdminGroup := v1Group.Group("/admin")
	{
		v1AdminGroup.GET("/users", func(c *gin.Context) {
			c.String(200, "/v1/admin/users")
		})
		v1AdminGroup.GET("/manager", func(c *gin.Context) {
			c.String(200, "/v1/admin/manager")
		})
		v1AdminGroup.GET("/photo", func(c *gin.Context) {
			c.String(200, "/v1/admin/photo")
		})
	}

	//V2版本的API
	v2Group := r.Group("/v2")
	v2Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v2/users")
	})
	v2Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v2/products")
	})

	r.Run(":8081")
}
