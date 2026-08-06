package main

import (
	"fmt"
	"log"
	"time"
)

func toque(msg string, tempo int) {

	for {
		log.Println(msg)
		time.Sleep(time.Second + time.Duration(tempo))
	}
}

func main() {

	go toque("tum", 1)
	go toque("ta", 2)

	stop := ""
	fmt.Scanf("%s", &stop)

}
