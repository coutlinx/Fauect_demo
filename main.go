package main

import (
	"fmt"
	S "linx/Final_Project/Hadnder"
	"time"
)

func main() {
	fmt.Println(time.Now())
	err :=S.Start()
	if err != nil{
		panic(err)
	}

}






