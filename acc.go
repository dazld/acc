package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var client = &http.Client{
	Timeout: time.Second * 1,
}

type Nested struct {
	Headers struct {
		Accept string `json:"Accept"`
	} `json:"headers"`
}

func main() {

	r, err := client.Get("https://httpbin.org/headers")

	if err != nil {
		spew.Dump(err)
		return
	}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		spew.Dump(err)
		return
	}
	var heads Nested
	if err := json.NewDecoder(r.Body).Decode(&heads); err != nil {
		spew.Dump(err)
	}
	spew.Dump(buf)
	spew.Dump(heads.Headers.Accept)
}
