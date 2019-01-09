package main

import (
	"encoding/json"
	"config/model"
	"config/route"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

// 配置文件路径
var configPath = flag.String("cp", "config.json", "config path")

func main() {
	port := flag.String("port", ":8080", "http listen port")
	flag.Parse()
	fmt.Println("配置文件为nil")
	file, err := os.Open(*configPath)
	if err != nil {
		panic("配置文件读取出错:" + err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := model.Configuration{}
	if err = decoder.Decode(&conf); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%#v", conf)
	if conf.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	}
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	model.GetDbConfig(&conf)
	r := route.GetHttpRouter() //获得路由实例
	tmpPort := *port
	log.Info("====管理员管理系统 启动  端口 ： " + tmpPort + " ====")
	r.Run(tmpPort)
}
