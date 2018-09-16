package main

import "fmt"

func main(){
	a := 1

	if b := 2; a == 1 && b ==2 {
		fmt.Println("Beethoven")
	} else if a == 1{
		fmt.Println("Mozart")
	} else {
		fmt.Println("Handel")
	}

	switch a {
	case 0 : 
		fmt.Println("Napoleon")
	case 1 : 
		fmt.Println("Vo Nguyen Giap")
	case 2 :
		fmt.Println("Zhukov")
	}
}