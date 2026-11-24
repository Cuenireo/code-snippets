// 查星期几：输入日期，输出它是星期几。
//
// 运行：go run weekday.go [YYYY-MM-DD，默认今天]
package main

import (
	"fmt"
	"os"
	"time"
)

var weekdays = []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}

func main() {
	var d time.Time
	var err error
	if len(os.Args) > 1 {
		d, err = time.Parse("2006-01-02", os.Args[1])
		if err != nil {
			fmt.Println("日期格式错误，用 YYYY-MM-DD")
			os.Exit(1)
		}
	} else {
		d = time.Now()
	}
	fmt.Printf("%s 是%s\n", d.Format("2006-01-02"), weekdays[d.Weekday()])
}
