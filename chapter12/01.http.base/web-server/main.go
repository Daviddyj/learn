package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`hello: `))
	}))
	m.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("rank"))

	}))
	m.Handle("/history/xiaoqiang", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`小强老师的黑历史`))
	}))
	m.Handle("/history/daiyijie", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`戴一杰帅哥的黑历史`))
	}))
	m.Handle("/history", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		qp := request.URL.Query()
		name := qp.Get("name")
		writer.Write([]byte(fmt.Sprintf("%s,%s帅哥的黑历史", request.Method, name)))
	}))

	http.ListenAndServe(":8080", m)

	//err := http.ListenAndServe("127.0.0.1:8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	//	//time.Sleep(2 * time.Second)
	//	//qp := request.URL.Query()
	//	//data, _ := json.Marshal(qp)
	//	if request.Body == nil {
	//		writer.Write([]byte("no body"))
	//		return
	//	}
	//
	//	//data, _ = ioutil.ReadAll(request.Body)
	//	//defer request.Body.Close()
	//	writer.Write([]byte("hello daiyijie is cool"))
	//}))
	//if err != nil {
	//	log.Println(err)
	//}
}
