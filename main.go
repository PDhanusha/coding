package main

import (
	supermarket "assign/first"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type items struct {
	Apple  float64
	Orange float64
	Pears  float64
}

type item1 struct {
	Fruit []items
}

type item []items

var prices = item{
	{
		Apple:  20.1,
		Orange: 24.4,
		Pears:  27.9,
	},
}

func main() {
	//new1 := new(items) //creating object
	supermarket.Get("apple")
	supermarket.Post("orange", 29.7)
	supermarket.Update("apple", 21.1)
	supermarket.Delete("orange")

	http.HandleFunc("/testget", testget) // each request calls handler
	http.HandleFunc("/testpost", testpost)
	http.HandleFunc("/testput", testput)
	//http.HandleFunc("/testdelete", testdelete)
	log.Fatal(http.ListenAndServe(":9455", nil))
}

func process(w http.ResponseWriter, cr chan *http.Request) {
	r := <-cr
	var s items
	json.NewDecoder(r.Body).Decode(&s)
	json.NewEncoder(w).Encode(s)

}

//l
func testget(w http.ResponseWriter, r *http.Request) {
	cr := make(chan *http.Request, 1)
	cr <- r
	var pleasewait sync.WaitGroup
	pleasewait.Add(1)

	go func() {
		defer pleasewait.Done()
		process(w, cr)
	}()
	pleasewait.Wait()
	w.WriteHeader(200)
}

func testpost(w http.ResponseWriter, r *http.Request) {
	var newitem items
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter item and value to create")
	}
	json.Unmarshal(reqBody, &newitem)
	prices = append(prices, newitem)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newitem)
}
func testput(w http.ResponseWriter, r *http.Request) {
	var newitem items
	err := json.NewDecoder(r.Body).Decode(&newitem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Println("Product Info - Updated")
		fmt.Println("apple ", newitem.Apple)
		fmt.Println("orange: ", newitem.Orange)
		fmt.Println("pears: ", newitem.Pears)
		w.WriteHeader(http.StatusOK)
		result, _ := json.Marshal(newitem)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

//l
// func (sel *items) testdelete(w http.ResponseWriter, r *http.Request, item string) {
// 	//var newitem item1.Fruit
// 	// vars := mux.Vars(r)
// 	// item := vars["item"]
// 	temp := items{Apple: 20.4, Orange: 29.7, Pears: 24.1}
// 	js, err := json.Marshal(temp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for index, prices := range items {
// 		if prices == item {
// 			sel.Fruit = append(sel.Fruit[:index], sel.Fruit[index+1:]...)
// 		}
// 	}
// 	w.Write(js)
// }
