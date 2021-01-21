package main

import (
	"bubble/routers"
	"fmt"
	"os"
)

func main() {
	port := os.Args[1]
	if len(port) < 1 {
		port = "18080"
	}
	port = ":" + port
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(port); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
