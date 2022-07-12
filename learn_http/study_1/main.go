package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://192.168.1.188:10000/open-api-service/region/v1?type=DB&ip=14.215.177.38")
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(body))
}
