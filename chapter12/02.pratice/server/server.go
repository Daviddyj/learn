package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learn/chapter12/02.pratice/frinterface"
	"learn/pkg/apis"
<<<<<<< HEAD
=======
	"log"
>>>>>>> 0efc8eb... 0327
	"net/http"
	"strings"
)

func main() {
	var rankServer frinterface.ServerInterface = NewFatRateRank()
	m := http.NewServeMux()
	m.Handle("/register", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "post") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if request.Body == nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
		payload, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法读取数据%s:", err)))
			return
		}
		var pi *apis.PersonInformation
		if err = json.Unmarshal(payload, &pi); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据%s:", err)))
			return
		}
		if err = rankServer.RegisterPersonInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("服务器注册用户失败:%s:", err)))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("success"))

	}))

	m.Handle("/personalinfo", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "post") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if request.Body == nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
		payload, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法读取数据%s:", err)))
			return
		}
		var pi *apis.PersonInformation
		if err = json.Unmarshal(payload, &pi); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据%s:", err)))
			return
		}
		if ft, err := rankServer.UpdatePersonInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("服务器注册用户失败:%s:", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(ft)
			writer.Write(data)
		}

	}))

	m.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "get") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		name := request.URL.Query().Get("name")
		if name == "" {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("参数未设置"))
			return
		}
		if rank, err := rankServer.GetFateRate(name); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("服务器获取用户排行榜失败:%s:", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(rank)
			writer.Write(data)
		}
	}))

	m.Handle("/ranktop", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "get") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if frTop, err := rankServer.GetTop(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("服务器获取用户排行榜失败:%s:", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(frTop)
			writer.Write(data)
		}
	}))

<<<<<<< HEAD
	http.ListenAndServe(":8080", m)
=======
	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Fatal(err)
	}
>>>>>>> 0efc8eb... 0327
}
