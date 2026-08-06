package main

import (
	"fmt"
	"log"
	"time"
)

func dizAlgo(msg string) {
	for{
		log.Println(msg)
		time.Sleep(time.Second + 1)
	}
}

func main() {
	go dizAlgo("Oi")
	go dizAlgo("Tchau")

	// Digitar enter para finalizar
	var ok string
	fmt.Scanf("%d", &ok)

	log.Println("Fim")
}
