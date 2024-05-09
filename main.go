package main

import "fmt"

func add(b []int, c chan int) {
	sum := 0
	for _, v := range b {
		sum += v
	}
	c <- sum
}

func main() {
	b := []int{3, 4, 3, 1, -5, 2}
	c := make(chan int)
	go add(b[:len(b)/2], c)
	go add(b[len(b)/2:], c)


	x, y := <-c, <-c

  //close channel will not take no more value to channnel
  close(c)

  // ok would be false if there is no value to retrieve or the channel is closed
  z,ok:= <-c
  fmt.Println("check",z,ok)

	fmt.Printf("x: %v, y:%v, x+y:%v\n", x, y, x+y)
}
