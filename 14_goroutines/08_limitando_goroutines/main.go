package main

import (
	"fmt"
	"sync"
)

func worker(id int, trabalhos <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for trabalho := range trabalhos {
		fmt.Printf("worker %d processou %d\n", id, trabalho)
	}
}

func main() {
	const quantidadeWorkers = 3
	trabalhos := make(chan int)

	var wg sync.WaitGroup
	for id := 1; id <= quantidadeWorkers; id++ {
		wg.Add(1)
		go worker(id, trabalhos, &wg)
	}

	for trabalho := 1; trabalho <= 10; trabalho++ {
		trabalhos <- trabalho
	}
	close(trabalhos)

	wg.Wait()
}