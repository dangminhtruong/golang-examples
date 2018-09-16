package main

import "fmt"

func main(){
	s := sum(1.2, 8.8)
	_, p  := split(10, 85)
	
	fmt.Println(s)
	fmt.Println(p)
}

func sum(a, b float64) float64 {
	return a + b
}

func split(a, b float64) (float64, float64) {
	return a / 2, b / 5
}