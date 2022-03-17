package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learn/pkg/apis"
	"log"
	"net/http"
	"time"
)

type frClient struct {
	handRing Interface
}

func (f frClient) register(pi apis.PersonInformation) {
	data, _ := json.Marshal(pi)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://127.0.0.1:8080/register", "application/json", r)
	if err != nil {
		log.Println("WARNING:register fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("返回:", string(data))
	}
}

func (f frClient) update(pi apis.PersonInformation) {
	data, _ := json.Marshal(pi)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://127.0.0.1:8080/personalinfo", "application/json", r)
	if err != nil {
		log.Println("WARNING:update fails:", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("返回:", string(data))
	}
}

func main() {
	fakePersonInformation := &randPersonInformation{}
	pi, _ := fakePersonInformation.ReadPersonInformation()
	frCli := frClient{}
	frCli.register(pi)
	for {
		pi, _ := fakePersonInformation.ReadPersonInformation()
		frCli.update(pi)
		time.Sleep(1 * time.Second)
	}
}
