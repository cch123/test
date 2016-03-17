package main

import "github.com/gin-gonic/gin"
import "net/http"

//import "reflect"

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "success",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/ping", handler)

	r.GET("/user/:name", func(c *gin.Context) {
		//取到的都是string类型，如果需要int需要自己转换
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	//注意这个例子和上面那个的区别
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.Run() // listen and server on 0.0.0.0:8080
}
