package main

import "fmt"


var m map[string]string

func main() {
	// array

	var arr [2]int
	arr[0] = 5
	arr[1] = 10

	for _,val := range arr {
		fmt.Println(val)
	}

	// Slice
	slices := []int{1, 2, 3}
	for _,slice := range slices {
		fmt.Println(slice)
	}

	// Map
	m = make(map[string]string)
	m["Name"] = "Dang Minh Truong"
	fmt.Println(m["Name"])

	n := map[string]string{"Address" : "Quang Binh"}
	fmt.Println(n["Address"])

	
}
