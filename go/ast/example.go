package lp

type person struct {
	age, aaa int `json:"age" form:"age"`
	name     int `json:"name" form:"name"`
	info     struct {
		detail int `json:"detail"`
	}
}
