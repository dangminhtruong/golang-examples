package main

import "fmt"

func main(){
	s := summary(1.2, 8.8)
	_, p  := spliter(10, 85)
	
	fmt.Println(s)
	fmt.Println(p)

	j,k := nakedReturn(10)

	fmt.Println(j)
	fmt.Println(k)
}

func summary(a, b float64) float64 {
	return a + b
}

func spliter(a, b float64) (float64, float64) {
	return a / 2, b / 5
}

func nakedReturn(a int) (m, n int){
	m = a / 2
	n = a / 5

	return
}