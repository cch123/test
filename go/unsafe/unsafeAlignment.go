package main

import (
	"fmt"
	"unsafe"
)

type V struct {
	i int32
	j int64
}

func (this *V) PutI() {
	fmt.Printf("i=%d\n", this.i)
	fmt.Printf("i=%v\n", &(this.i))
}

func (this *V) PutJ() {
	fmt.Printf("j=%d\n", this.j)
	fmt.Printf("j=%v\n", &(this.j))
}

func main() {
	var v *V = new(V)
	var i *int32 = (*int32)(unsafe.Pointer(v))
	*i = int32(98)
	fmt.Println(unsafe.Pointer(v))
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int64(0)))))
	//var j = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int64(0)))))
	var j = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0)))))
	*j = int64(763)
	v.PutI()
	v.PutJ()
	fmt.Println(unsafe.Sizeof(*v))
	fmt.Println(unsafe.Sizeof((*v).i))

}
