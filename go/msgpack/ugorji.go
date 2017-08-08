// 注意，mspack 编码的字段必须是大写字母开头(可导出
// 内部应该和 json 一样是用反射操作的
package main

import (
	"fmt"

	"github.com/ugorji/go/codec"
	"github.com/vmihailenco/msgpack"
)

func main() {
	type person struct {
		age    int
		Name   string
		Height int
	}
	var v1 = person{11, "alex", 2323}
	var b []byte = make([]byte, 0, 64)
	var h codec.Handle = new(codec.MsgpackHandle)
	var enc *codec.Encoder = codec.NewEncoderBytes(&b, h)
	var err error = enc.Encode(v1) //any of v1 ... v8
	// b now contains the encoded value.
	fmt.Println(err)
	fmt.Println(string(b))
	fmt.Println(len(b))
	res, _ := msgpack.Marshal(v1)
	fmt.Println(len(res))
}
