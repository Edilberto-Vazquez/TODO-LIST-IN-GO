package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string, c chan<- string) {
	// defer wg.Done()
	_, err := http.Get(server)
	if err != nil {
		c <- "the server " + server + " is down =("
	} else {
		c <- "the server " + server + " is ok =)"
	}
}

func main() {
	start := time.Now()
	// var wg sync.WaitGroup
	c := make(chan string, 5)
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://twitch.com",
		"http://instagram.com",
		"http://twitter.com",
	}

	i := 0

	for i < 2 {
		for _, v := range servidores {
			// wg.Add(1)
			go checkServer(v, c)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-c)
		i++
	}

	// wg.Wait()
	// for i := 0; i < len(servidores); i++ {
	// 	fmt.Println(<-c)
	// }

	end := time.Since(start)
	fmt.Printf("execution time %s\n", end)
}
