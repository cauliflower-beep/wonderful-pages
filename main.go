package main

import (
	"fmt"
	"os"
	"wonderful-pages/routers"
	"wonderful-pages/setting"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./conf/config.ini")
		return
	}
	// 加载配置文件 os.Args[1]为配置文件路径
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
	fmt.Println("server is listening ...")
}
