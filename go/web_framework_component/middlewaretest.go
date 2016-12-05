package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func main() {
	signature := ""
	router := gin.New()
	//router.HandleMethodNotAllowed = true
	router.Use(func(c *gin.Context) {
		println("entering 1")
		signature += "A"
		c.Next()
		signature += "B"
		println("leaving 1")
	})
	router.Use(func(c *gin.Context) {
		println("entering 2")
		signature += "C"
		c.Next()
		signature += "D"
		println("leaving 2")
	})
	router.Use(func(c *gin.Context) {
		println("entering 3")
		signature += "E"
		c.Next()
		signature += "F"
		println("leaving 3")
	}, func(c *gin.Context) {
		println("entering 4")
		signature += "G"
		c.Next()
		signature += "H"
		println("leaving 4")
	})
	router.GET("/", func(c *gin.Context) {
		println("entering main")
		signature += " fuck "
		println("leaving main")
	})
	router.GET("/get", func(c *gin.Context) {
		println("entering main")
		signature += " xxx"
		println("leaving main")
	})
	// RUN
	w := performRequest(router, "GET", "/")
	fmt.Println(w.Code)
	fmt.Println(signature)

	w = performRequest(router, "GET", "/get")
	fmt.Println(w.Code)
	fmt.Println(signature)
}
