package main

import "fmt"

type iosDeveloper struct {
	languge string
	canUseXcode bool
}

type androidDevelop struct {
	languge string
	rank float64
}

type developer interface {
	canCodeAnLanguge() string
}

func main() {
	anhDoanh := androidDevelop{
		"Kotlin",
		5.0,
	}

	anhHoang := iosDeveloper{
		"Swift",
		true,
	}

	showDeveloperSkill(anhDoanh)
	showDeveloperSkill(anhHoang)
}


func showDeveloperSkill(dev developer) {
	fmt.Println(dev.canCodeAnLanguge())
}

func (iosDev iosDeveloper) canCodeAnLanguge() string {
	return iosDev.languge
}

func (andDev androidDevelop) canCodeAnLanguge() string {
	return andDev.languge
}