package main

import "fmt"

type person struct{
	firtsName string
	lasname string
	contac contacs
}

type contacs struct {
	email string
	phone string
}

func main(){
	me := person{
		"Dang Minh",
		"Truong",
		contacs {
			"xdangminhtruongx@gmail.com",
			"0965296242",
		},
	}
	me.updateFirstName("Dang")
	me.updateEmail("dang.minh.truong@framgia")
	me.showAllInfor();
}

func (p *person) updateFirstName(newFirstName string){
	p.firtsName = newFirstName
}


func (p person) showAllInfor() {
	fmt.Println(p);
}

func (p *person) updateEmail(newEmail string) {
	p.contac.email = newEmail
}
