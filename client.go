package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type items struct {
	Apple  float64
	Orange float64
	Pears  float64
}

func main() {
	client := &http.Client{}
	client.Timeout = time.Second * 15
	temp := items{Apple: 20.4, Orange: 29.7, Pears: 24.1}
	j, err := json.Marshal(temp)
	req, err := http.NewRequest("GET", "http://localhost:9455/testget", bytes.NewBuffer(j))
	if err != nil {
		fmt.Println("1st err")
		panic(err)
	}
	//defer req.Body.Close()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("2nd err")
		panic(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("3rd err")
		panic(err)
	}
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Body : ", string(respBody))
}
