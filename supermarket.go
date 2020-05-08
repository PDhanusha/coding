package supermarket

import (
	"fmt"
)

//l
func Get(item string) {
	price := map[string]float64{"apple": 20.4, "pears": 24.1}
	value, ok := price[item]
	if ok {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("key not found")
	}
}

//l
func Post(item string, val float64) {
	price := map[string]float64{"apple": 20.4, "pears": 24.1}
	value, ok := price[item]
	if ok {
		fmt.Println("key already exist")
	} else {
		price[item] = val
		fmt.Println(price)
		fmt.Println(value) // showing error "value declared and not used" if not
	}
}

//u
func Update(item string, val float64) {
	price := map[string]float64{"apple": 20.4, "orange": 29.7, "pears": 24.1}
	value, ok := price[item]
	if ok {
		fmt.Println("key already exist")
		price[item] = val
		fmt.Println(price)
		fmt.Println(value)
	} else {
		fmt.Println("key does not exist")
	}
}

//k
func Delete(item string) {
	price := map[string]float64{"apple": 20.4, "orange": 29.7, "pears": 24.1}
	delete(price, item)
	fmt.Println(price)
}
