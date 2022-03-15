package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	err := http.ListenAndServe("127.0.0.1:8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(2 * time.Second)
		writer.Write([]byte("hello daiyijie is cool"))
	}))
	if err != nil {
		log.Println(err)
	}
}
