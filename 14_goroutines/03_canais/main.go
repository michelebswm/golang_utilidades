package main

import (
	"fmt"
	"log"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
		time.Sleep(time.Second * time.Duration(1))
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
		time.Sleep(time.Second * time.Duration(1))
	}
}

func print(c chan string) {
	for {
		msg := <-c
		log.Println(msg)
	}
}

func main() {
	// Uma forma de criar um canal
	// var canal1 chan string = make(chan string)

	canal := make(chan string)
	log.Println(canal) // Retorna um endereço

	go ping(canal)
	go pong(canal)
	go print(canal)

	ok := ""
	fmt.Scanf("$s", ok)

}
