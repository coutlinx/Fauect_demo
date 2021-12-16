package main

import (
	"fmt"
	"github.com/robfig/cron"
	C "linx/Final_Project/Config"
	S "linx/Final_Project/Hadnder"
	"time"
)

func main() {
	c := cron.New() // 新建一个定时任务对象
	err := c.AddFunc("0 0 1 * * ?", C.FreeLimits)
	if err != nil {
		return
	}
	c.Start()
	select {
	default:
		fmt.Println(time.Now())
		err = S.Start()
		if err != nil {
			panic(err)
		}

	} // 给对象增加定时任务
	//c := cron.New()
	//err = c.AddFunc("*/5 * * * * ?",C.FreeLimits)
	//if err != nil{
	//	panic(err)
	//}
	//c.Start()
	//select {
	//}
}
