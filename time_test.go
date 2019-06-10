package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	//获取本地location
	toBeCharge := "2015-01-01 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板，必须用这个模版据说是因为是golang诞生的日期
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println("theTime ", theTime)                                //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println("sr ", sr)                                          //打印输出时间戳 1420041600

	//时间戳转日期
	dataTimeStr := time.Unix(1549940849, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println("dataTimeStr ", dataTimeStr)
}

func TestCompare(t *testing.T) {

	//返回现在时间
	tNow := time.Now()
	//时间转化为string，layout必须为 "2006-01-02 15:04:05"
	timeNow := tNow.Format("2006-01-02 15:04:05")
	fmt.Println("tNow(time format):", tNow)
	fmt.Println("tNow(string format):", timeNow)

	//string转化为时间，layout必须为 "2006-01-02 15:04:05"
	time, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	fmt.Println("t(time format)", time)

	//某个时间点 前后判断
	trueOrFalse := time.After(tNow)
	if trueOrFalse == true {
		fmt.Println("t（2014-06-15 08:37:18）在tNow之后!")
	} else {
		fmt.Println("t（2014-06-15 08:37:18）在tNow之前!")
	}
	fmt.Println()
}

func TestTimeNow(t *testing.T) {
	n := time.Now().Format("2006-01-02 15:04:05")
	t.Log(n)
}

// 多时区
func TestMultipleTimeZones(t *testing.T) {
	now := time.Now()
	local1, err1 := time.LoadLocation("") // 等同于"UTC"
	if err1 != nil {
		fmt.Println(err1)
	}
	local2, err2 := time.LoadLocation("Local") // 服务器设置的时区
	if err2 != nil {
		fmt.Println(err2)
	}
	local3, err3 := time.LoadLocation("America/Los_Angeles") // PDT
	if err3 != nil {
		fmt.Println(err3)
	}
	// 向东每个时区+1，向西每个时区-1
	fmt.Println(now.In(local1)) // UTC(Universal Time Coordinated世界协调时间)，又叫做0时区
	fmt.Println(now.In(local2)) // CST(China Standard Time北京时间)，又叫做东8区
	fmt.Println(now.In(local3)) // PDT（美国时区），又叫做西8区
}

// 比较时间大小
func TestTimeCompare(t *testing.T) {
	now := time.Now()
	// 今天10点
	todayClock := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, now.Location())
	// 当前时间在今天10点之后
	b := now.After(todayClock)
	fmt.Println(b)
	// 当前时间在今天10点之前
	b = now.Before(todayClock)
	fmt.Println(b)
	return
}
