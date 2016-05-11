package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 获取大小的借口
type Sizer interface {
	Size() int64
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
	println("hello")
	if "POST" == r.Method {
		file, _, err := r.FormFile("userfile")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer file.Close()
		f, err := os.Create("filenametosaveas")
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintf(w, "上传文件的大小为: %d", file.(Sizer).Size())
		return
	}

	// 上传页面
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	html := `
<form enctype="multipart/form-data" action="/hello" method="POST">
    Send this file: <input name="userfile" type="file" />
    <input type="submit" value="Send File" />
</form>
`
	io.WriteString(w, html)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
