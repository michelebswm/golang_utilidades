package main

import (
	"fmt"
	"os"
)

func main() {
	valor, existe := os.LookupEnv("APP_ENV")
	if !existe {
		fmt.Println("APP_ENV não foi definida")
		return
	}

	fmt.Println("Ambiente:", valor)
}
