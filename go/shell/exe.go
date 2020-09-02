package main
import (
	"bytes"
	"fmt"
	"errors"
	"os/exec"
)

func main() {
	cmd := exec.Command("/Users/xargin/go/bin/go", "build", "fuck")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Println("err is ", err)
	if err != nil {
		println("fuck")
		var ee *exec.ExitError
		if errors.As(err, &ee)  {
			e := err.(*exec.ExitError)
			fmt.Println(e.ExitCode())
		}
	}
}
