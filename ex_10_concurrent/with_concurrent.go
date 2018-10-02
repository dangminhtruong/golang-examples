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

	c := make(chan string)
	start := time.Now()

	for _, link :=  range links{
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println( <- c )
	}

	ended := time.Since(start)
    fmt.Println("All request took %s", ended)
}

func checkLink(link string, c chan string){
	_, err := http.Get(link)

	if err != nil {
		c <- link + " Maybe down!"
		return 
	}
	
	c <- link + " is up and running!"
} 
