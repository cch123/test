package main

import (
	"reflect"

	"github.com/gin-gonic/gin"
)
import "net/http"
import "os"
import "io"
import "fmt"
import "log"

//import "reflect"

func handler(c *gin.Context) {
	var a interface{}
	err := c.BindJSON(&a)
	if err != nil {
	}

	fmt.Println(reflect.TypeOf(a.(map[string]interface{})["a"]))
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "success",
	})
}

func main() {
	r := gin.Default()

	//router用法
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

	/*
		//注意这个例子和上面那个的区别
		r.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
	*/

	//url里有querystring的
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	//表单post和有默认值的表单post
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	//既有form又有querystring的
	//POST /post?id=1234&page=1 HTTP/1.1
	//Content-Type: application/x-www-form-urlencoded

	//name=manu&message=this_is_great
	r.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	//post文件上传
	r.POST("/upload", func(c *gin.Context) {

		file, header, err := c.Request.FormFile("upload")
		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create("./tmp/" + filename + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	})

	// 路由分组，用来做api的版本管理
	// Simple group: v1
	v1 := r.Group("/v1")
	{
		v1.POST("/login", handler)
		v1.POST("/submit", handler)
		v1.POST("/read", handler)
	}

	// Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", handler)
		v2.POST("/submit", handler)
		v2.POST("/read", handler)
	}
	r.GET("/post/:id", func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		fmt.Println("post id")
	})

	r.GET("/post/:id/reply/add", func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		fmt.Println("post id reply add")
	})
	r.POST("/post/add", func(c *gin.Context) {
		fmt.Println("post add")
	})

	r.Run() // listen and server on 0.0.0.0:8080

}
