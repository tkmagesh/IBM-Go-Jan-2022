package main

import (
	"fmt"
)

//Share memory by communicating (using channels)
func main() {

	ch := make(chan int)
	go writeData(ch)
	result := <-ch
	fmt.Println(result)

}

func writeData(ch chan int) {
	ch <- 100
}
