package main

import "github.com/gin-gonic/gin"

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	r := gin.Default()

	r.GET("/users/123", func(c *gin.Context) {
		//c.JSON(200, user{ID: 123, Name: "张三", Age: 20})

		//美化JSON
		c.IndentedJSON(200, user{ID: 123, Name: "张三", Age: 20})
	})

	_ = r.Run(":8081")

}
