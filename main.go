package main

import (
	"bubble/routers"
	"fmt"
)

func main() {
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(":18080"); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
