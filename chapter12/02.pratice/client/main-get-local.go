package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
<<<<<<< HEAD
	"learn/pkg/apis"
=======
>>>>>>> 0efc8eb... 0327
	"log"
	"net/http"
	"time"
)

type frClient struct {
	handRing Interface
}

<<<<<<< HEAD
func (f frClient) register(pi apis.PersonInformation) {
=======
func (f frClient) register() {
	pi, _ := f.handRing.ReadPersonInformation()
>>>>>>> 0efc8eb... 0327
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

<<<<<<< HEAD
func (f frClient) update(pi apis.PersonInformation) {
=======
func (f frClient) update() {
	pi, _ := f.handRing.ReadPersonInformation()
>>>>>>> 0efc8eb... 0327
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
<<<<<<< HEAD
	fakePersonInformation := &randPersonInformation{}
	pi, _ := fakePersonInformation.ReadPersonInformation()
	frCli := frClient{}
	frCli.register(pi)
	for {
		pi, _ := fakePersonInformation.ReadPersonInformation()
		frCli.update(pi)
=======
	//fakePersonInformation := &randPersonInformation{}

	frCli := frClient{handRing: &randPersonInformation{}}
	frCli.register()
	for {
		frCli.update()
>>>>>>> 0efc8eb... 0327
		time.Sleep(1 * time.Second)
	}
}
