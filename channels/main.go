package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://www.google.com", "http://facebook.com", "http://stackoverflow.com", "http://amazon.com", "http://golang.org", "http://fake.url.com"}

	c := make(chan string)

	for _, l := range links {
		go checkStatus(l, c)
	}

	for link := range c {
		go func(l string) {
			time.Sleep(time.Second)
			checkStatus(l, c)
		}(link)
	}
}

func checkStatus(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, "might be down.")
		c <- l
		return
	}

	fmt.Println(l, "is up!")
	c <- l
}
