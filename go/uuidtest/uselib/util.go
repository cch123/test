package main

import uuid "github.com/satori/go.uuid"

var u1 uuid.UUID
var namespaceDNS, _ = uuid.FromString("ss")

func test() {
	u1 = uuid.NewV4()
	uuid.NewV5(namespaceDNS, "www.example.com")
}
