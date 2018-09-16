package main

import(
	"fmt"
	"net/http"
	"time"
)

func main(){
	links := []string{
		"https://www.google.com.vn",
		"https://stackoverflow.com",
		"https://facebook.com",
		"https://www.amazon.com/",
		"https://golang.org/",
		"https://viblo.asia/",
		"https://medium.com/",
		"https://www.youtube.com/",
		"https://github.com/",
		"https://www.digitalocean.com/",
	}

    start := time.Now()
	for _, link :=  range links{
		checkLink(link)
	}

	ended := time.Since(start)
    fmt.Println("All request took %s", ended)
}

func checkLink(link string){
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " maybe down") 
	}
	
	fmt.Println(link + " is up") 
} 