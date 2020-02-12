package main

import (
	"fmt"
	"net/http"
)

func main() {
	httpGetter := HTTPGetter{client: &http.Client{}}
	resp, err := httpGetter.Get("http://worldtimeapi.org/api/ip")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
