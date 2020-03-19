package main

import (
	"fmt"
	"net/http"

	"github.com/reedthink/pkg/setting"
	"github.com/reedthink/routers"
)

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
