package main

import "net/http"

func main() {
	for {
		go func() {
			req, _ := http.NewRequest("GET", "www.baidu.com", nil)
			client := http.DefaultClient
			client.Do(req)
			// should be resp, _ := client.Do(req
			// defer resp.Body.Close
		}()
	}
}
