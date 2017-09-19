package main

import "fmt"

type IMessage interface {
	Print()
	OutPut()
}

type BaseMessage struct {
	Msg string
}

func (message *BaseMessage) Print() {
	fmt.Println("baseMessage:msg", message.Msg)
	message.OutPut()
}

func (message *BaseMessage) OutPut() {
	//fmt.Println(reflect.TypeOf(message))
	fmt.Println("baseMessage:out")
}

type SubMessage struct {
	BaseMessage
}

func (message *SubMessage) Print() {
	message.BaseMessage.Print()
}

func (message *SubMessage) OutPut() {
	fmt.Println("subMessage:out")
}

func interface_use(i IMessage) {
	i.Print()
}

func main() {
	baseMessage := new(BaseMessage)
	baseMessage.Msg = "a"
	interface_use(baseMessage)

	subMessage := &SubMessage{BaseMessage{Msg: "aaa"}}
	subMessage.Msg = "b"
	interface_use(subMessage)
}
