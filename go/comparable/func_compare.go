package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var myHandler struct {
	handler  func(*gin.Context)
	fullPath string
}

func initMyHandler(fullPath string) {
	myHandler.handler = func(ctx *gin.Context) {
		fmt.Println(ctx.Param("id"))
	}
	myHandler.fullPath = fullPath
}

func main() {
}
