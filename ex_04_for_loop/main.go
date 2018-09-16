package main

import "fmt"

func main(){
	forLoopSimple()
}


func forLoopSimple() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forWhile() {
	sum := 1
	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)
}

func foreverForFun(){
	for {
	}
}