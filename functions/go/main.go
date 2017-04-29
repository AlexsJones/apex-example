package main

import (
	"encoding/json"

	"github.com/apex/go-apex"
)

type message struct {
	Name string `json:"name"`
}

type response struct {
	Reply string `json:"reply"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message

		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}
		r := response{"Yo, " + m.Name}
		return r, nil
	})
}
