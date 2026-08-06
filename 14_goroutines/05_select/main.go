package main

import (
	"fmt"
	"log"
	"time"
)

func ping(c chan string) {
	for i := 1; i <= 10; i++ {
		c <- "ping"
		time.Sleep(time.Second * time.Duration(1))
	}
}

func pong(c chan string) {
	for i := 1; i <= 15; i++ {
		c <- "pong"
		time.Sleep(time.Second * time.Duration(1))
	}
}

func print(c1, c2 chan string) {
	for {
		select {
		case msg := <-c1:
			log.Println("Canal 1: ", msg)

		case msg := <-c2:
			log.Println("Canal 2: ", msg)

		case <-time.After(time.Second * time.Duration(2)):
			fmt.Println("Tempo esgotado")
			return
		}

	}
}

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go ping(canal1)
	go pong(canal2)
	print(canal1, canal2)

}
