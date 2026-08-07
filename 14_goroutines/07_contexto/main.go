package main

import (
	"context"
	"fmt"
	"time"
)

func trabalhar(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("trabalho cancelado:", ctx.Err())
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Println("trabalhando...")
		}
	}
}

func main() {
	ctx, cancelar := context.WithTimeout(context.Background(), 700*time.Millisecond)
	defer cancelar()

	trabalhar(ctx)
}