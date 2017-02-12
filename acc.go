package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var client = &http.Client{
	Timeout: time.Second * 1,
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
	spew.Dump(buf)
}
