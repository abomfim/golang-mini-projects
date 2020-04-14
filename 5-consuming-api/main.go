package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiurl := "http://viacep.com.br/ws/01001000/json"
	response, err := http.Get(apiurl)

	if err != nil {
		panic(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseData))
}
