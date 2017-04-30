package main

import (
	"encoding/json"
	"strconv"

	"github.com/apex/go-apex"
)

type message struct {
	Count int `json:"count"`
}

type response struct {
	Results []string `json:"results"`
}

func fibGenerator() chan int {
	c := make(chan int)

	go func() {
		for i, j := 0, 1; ; i, j = i+j, i {
			c <- i
		}
	}()

	return c
}
func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message

		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}
		c := m.Count
		c1 := fibGenerator()

		results := make([]string, c)
		for x := 0; x < c; x++ {
			answer := strconv.Itoa(<-c1)
			results = append(results, answer)
		}

		r := response{Results: results}

		return r, nil
	})
}
