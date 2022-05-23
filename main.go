package main

import (
	"douyin/routers"
	"fmt"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(":8000"); err != nil{
		fmt.Println("Start service failled, err:%v\n", err)
	}
}
