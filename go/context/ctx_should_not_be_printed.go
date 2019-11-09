package main

import (
	"fmt"
	"net/http"
)

func panic(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context())
}

func main() {
	http.HandleFunc("/", panic)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}

//wrk -c200 -t20 -d30s  http://127.0.0.1:9090
/*
package main

import (
    "fmt"
    "net/http"
    "reflect"
)

func panic(w http.ResponseWriter, r *http.Request) {
    server := r.Context().Value(http.ServerContextKey).(*http.Server)
    v := reflect.ValueOf(*server)

    for i := 0; i < v.NumField(); i++ {
        if name := v.Type().Field(i).Name; name != "activeConn" {
            continue
        }
        fmt.Println(v.Field(i))
    }
}

func main() {
    http.HandleFunc("/", panic)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        fmt.Println(err)
    }
}

wrk -c200 -t20 -d30s  http://127.0.0.1:9090
*/
