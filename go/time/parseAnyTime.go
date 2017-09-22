package main

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

func main() {
	// timeparse 提供的增强 parse 函数
	// 大部分常见的格式都可以正常解析
	d, err := dateparse.ParseAny("2015-01-02")
	fmt.Println(d, err)

	// 按照程序的 time.Local 自动填上时区信息
	// 类似 time.ParseInLocation("layout", timeStr, time.Local)
	d, err = dateparse.ParseLocal("2015-01-02")
	fmt.Println(d, err)

	// 需要显式提供 local 参数的一个函数
	d, err = dateparse.ParseIn("2014-04-03", time.Local)
	fmt.Println(d, err)

	// 可以通过名字来初始化一个 time.Location 对象
	beijingLocal, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(beijingLocal)
	otherLocal, _ := time.LoadLocation("Europe/Budapest")

	// 通过 ParseIn，类似于 time.ParseInLocation
	// 这种设置了时区解析的，如果字符串里没有时区信息
	// 会把时区信息补到解析后的 time.Time 对象里
	// 实际上并不一定是你想要的效果
	d, err = dateparse.ParseIn("2015-03-02", otherLocal)
	fmt.Println(d, err)

	// 已有时区信息的时间戳
	// 可以通过 time 对象的 In 来切换时区
	fmt.Println(d.In(beijingLocal))

	// 所有支持的格式，还是蛮方便的
	// 确认一下有没有性能隐患
	/*
		"May 8, 2009 5:57:51 PM",
		"Mon Jan  2 15:04:05 2006",
		"Mon Jan  2 15:04:05 MST 2006",
		"Mon Jan 02 15:04:05 -0700 2006",
		"Monday, 02-Jan-06 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"Tue, 11 Jul 2017 16:28:13 +0200 (CEST)",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"Mon Aug 10 15:44:11 UTC+0100 2015",
		"Fri Jul 03 2015 18:04:07 GMT+0100 (GMT Daylight Time)",
		"12 Feb 2006, 19:17",
		"2013-Feb-03",
		//   mm/dd/yy
		"3/31/2014",
		"03/31/2014",
		"08/21/71",
		"8/1/71",
		"4/8/2014 22:05",
		"04/08/2014 22:05",
		"04/2/2014 03:00:51",
		"8/8/1965 12:00:00 AM",
		"8/8/1965 01:00:01 PM",
		"8/8/1965 01:00 PM",
		"8/8/1965 1:00 PM",
		"8/8/1965 12:00 AM",
		"4/02/2014 03:00:51",
		"03/19/2012 10:11:59",
		"03/19/2012 10:11:59.3186369",
		// yyyy/mm/dd
		"2014/3/31",
		"2014/03/31",
		"2014/4/8 22:05",
		"2014/04/08 22:05",
		"2014/04/2 03:00:51",
		"2014/4/02 03:00:51",
		"2012/03/19 10:11:59",
		"2012/03/19 10:11:59.3186369",
		//   yyyy-mm-ddThh
		"2006-01-02T15:04:05+0000",
		"2009-08-12T22:15:09-07:00",
		"2009-08-12T22:15:09",
		"2009-08-12T22:15:09Z",
		//   yyyy-mm-dd hh:mm:ss
		"2014-04-26 17:24:37.3186369",
		"2012-08-03 18:31:59.257000000",
		"2014-04-26 17:24:37.123",
		"2013-04-01 22:43:22",
		"2014-12-16 06:20:00 UTC",
		"2014-12-16 06:20:00 GMT",
		"2014-04-26 05:24:37 PM",
		"2014-04-26 13:13:43 +0800",
		"2014-04-26 13:13:44 +09:00",
		"2012-08-03 18:31:59.257000000 +0000 UTC",
		"2015-09-30 18:48:56.35272715 +0000 UTC",
		"2015-02-18 00:12:00 +0000 GMT",
		"2015-02-18 00:12:00 +0000 UTC",
		"2017-07-19 03:21:51+00:00",
		"2014-04-26",
		"2014-04",
		"2014",
		"2014-05-11 08:20:13,787",
		//  yyyymmdd and similar
		"20140601",
		// unix seconds, ms
		"1332151919",
		"1384216367189",
	*/
	d, err = dateparse.ParseAny("2015-03-04")
	fmt.Println(d, err)

}
