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
	router.HandleMethodNotAllowed = true
	router.Use(func(c *gin.Context) {
		signature += "A"
		c.Next()
		signature += "B"
	})
	router.Use(func(c *gin.Context) {
		signature += "C"
		c.Next()
		signature += "D"
	})
	router.Use(func(c *gin.Context) {
		signature += "E"
		c.Next()
		signature += "F"
	}, func(c *gin.Context) {
		signature += "G"
		c.Next()
		signature += "H"
	})
	router.GET("/", func(c *gin.Context) {
		signature += " fuck "
	})
	// RUN
	w := performRequest(router, "GET", "/")

	fmt.Println(w.Code)
	fmt.Println(signature)
}
