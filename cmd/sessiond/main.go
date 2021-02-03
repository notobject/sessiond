package main

import (
	"flag"
	"log"

	"github.com/notobject/sessiond/api/web"
	"github.com/notobject/sessiond/internal/session"
	"github.com/notobject/sessiond/pkg/config"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "config", "session_dev.yaml", "config file")

	// 解析命令行参数
	flag.Parse()

	// 读取配置文件
	cfg, err := config.Reload(configPath)
	if err != nil {
		log.Fatal("load config failed: ", err)
	}

	if cfg.Debug {
		log.Println("run mode DEBUG")
	}

	// 启动rpc服务
	session.Start(cfg)

	// 启动Web服务
	web.Start(cfg)

	// 阻塞主线程
	select {}
}
