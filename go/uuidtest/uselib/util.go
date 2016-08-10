package main

import uuid "github.com/satori/go.uuid"

var u1 uuid.UUID

func test() {
	u1 = uuid.NewV4()
}
