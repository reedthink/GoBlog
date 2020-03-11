package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "您看到我了")
}

func main() {
	http.HandleFunc("/", sayHello)
	log.Println("启动了")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("List 8000")
	}
}
