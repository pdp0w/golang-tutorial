package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	// specific time
	specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println("specific time : ", specificTime)

	// parse time
	parseTime, _ := time.Parse("2006-01-02", "2020-05-01")
	fmt.Println(parseTime)

	// 00:00:00 UTC on jan 1, 1970

	now := time.Now()
	unixTime := now.Unix()

	fmt.Println(unixTime)

	t := time.Unix(unixTime, 0)
	fmt.Println(t.Format("2006-01-02"))

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2024-07-04T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)
}
