package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func main() {
	buffer := new(bytes.Buffer)
	router := gin.New()
	//router.Use(gin.RecoveryWithWriter(buffer))
	router.GET("/recovery", func(_ *gin.Context) {
		panic("Oupps, Houston, we have a problem")
	})
	w := performRequest(router, "GET", "/recovery")
	// TEST
	fmt.Println(w.Code, 500)
	fmt.Println(buffer.String(), "GET /recovery")
	fmt.Println(buffer.String(), "Oupps, Houston, we have a problem")
	fmt.Println(buffer.String(), "TestPanicInHandler")
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
