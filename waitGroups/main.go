package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}

	wg.Add(len(words))

	for i, w := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}

	wg.Wait()

	// Very bad solution because we do not know how long does it need to take the goroutine to run
	// time.Sleep(1 * time.Second)

	wg.Add(1)
	printSomething("Second string", &wg)
}
