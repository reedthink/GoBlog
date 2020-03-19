package main

import (
	"fmt"
	"net/http"

	"github.com/reedthink/pkg/setting"
	"github.com/reedthink/routers"
)

// @title 测试
// @version 0.0.1
// @description  测试
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
ALTER user 'root'@'localhost' IDENTIFIED BY 'sqldb'
*/
