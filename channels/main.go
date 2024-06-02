package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!!!")
		c <- link
		return
	}

	fmt.Println(link, " is up and running!")
	c <- link
}

//func checkLink(link string, c chan string) {
//
//	_, err := http.Get(link)
//
//	if err != nil {
//		fmt.Println(link, " might be down!!!")
//		c <- link + " might be down!!!"
//		return
//	}
//
//	fmt.Println(link, " is up and running!")
//	c <- link + " is up!"
//}
