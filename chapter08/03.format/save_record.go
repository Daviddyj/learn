package main

import (
	"encoding/json"
	"learn/pkg/apis"
	"log"
	"os"
)

type record struct {
	filePath string
}

func NewRecord(filePath string) *record {
	return &record{
		filePath: filePath,
	}
}

func (c *record) savePersonInformation(pi *apis.PersonInformation) error {
	data, err := json.Marshal(pi)
	if err != nil {
		log.Fatal(err)
	}
	return c.writeFileWithAppend(data)
}

func (c *record) writeFileWithAppend(data []byte) error {
	file, err := os.OpenFile(c.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(append(data, "\n"...))
	if err != nil {
		return err
	}
	return nil
}
