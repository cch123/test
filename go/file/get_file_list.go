package main
import "io/ioutil"

import "fmt"
func main() {
	//ioutil.ReadFile(filenamoe)
	filePathList, _:=ioutil.ReadDir("./")
	for _, p := range filePathList {
		fmt.Println(p.Name())
	}
}