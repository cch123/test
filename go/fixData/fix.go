package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./driver_list.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		resArr := strings.Split(line, " ")
		if len(resArr) <= 2 {
			continue
		}

	}
}
