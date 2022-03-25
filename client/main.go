package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Client starting...")
	time.Sleep(3 * time.Second)

	_, err := http.Get("http://server:8080/hello")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Success!")
	}
}
