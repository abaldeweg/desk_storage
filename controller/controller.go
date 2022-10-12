package controller

import (
	"encoding/json"
	"log"
)

type File struct {
    Body string `json:"body"`
}

type Msg struct {
    Msg string `json:"msg"`
}

func init() {
    log.SetPrefix("controller: ")
    log.SetFlags(0)
}

func UnmarshalJson(blob []byte, d *[]interface{}) *[]interface{} {
	if err := json.Unmarshal(blob, &d); err != nil {
		log.Fatal(err)
	}

    return d
}

func MarshalJson(data interface{}) []byte {
	d, err := json.Marshal(&data)
    if err != nil {
        log.Fatal(err)
    }

    return d
}
