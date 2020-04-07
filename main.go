package main

import (
	"fmt"
	"net/http"

	"GoBlog/pkg/setting"
	"GoBlog/routers"
)

// @title GoBlog
// @version 0.0.2
// @description GoBlog的API文档
// @BasePath /api/v1/
func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
/*
重新生成API文档
swag init 

再次编译
go build

后台运行程序
nohup ./GoBlog &

杀掉后台
pkill GoBlog
*/