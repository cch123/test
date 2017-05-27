package main

import "time"

type StarEvent struct {
	EventID      int       `db:"event_id"`
	DriverID     int       `db:"driver_id"`
	AutoGrabFlag int       `db:"auto_grab_flag"`
	OrderID      int       `db:"order_id"`
	Area         int       `db:"area"`
	RegisterArea int       `db:"register_area"`
	OrderType    int       `db:"order_type"`
	Sid          int       `db:"sid"`
	CarLevel     int       `db:"car_level"`
	Star         int       `db:"star"`
	TagList      string    `db:"tag_list"`
	Comment      string    `db:"comment"`
	Status       int       `db:"status"`
	CommentTime  int       `db:"comment_time"`
	ModifyTime   time.Time `db:"modify_time"`
	CreateTime   time.Time `db:"create_time"`
	Status1      int       `db:"status1"`
}

var a = []StarEvent{
	StarEvent{EventID: 10},
	StarEvent{EventID: 10},
	StarEvent{EventID: 10},
	StarEvent{EventID: 10},
	StarEvent{EventID: 10},
	StarEvent{EventID: 10},
}
var b interface{} = a

func asx() {
	_ = b.([]StarEvent)
}
