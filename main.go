package main

import "fmt"

func add(b []int, c chan int) {
	sum := 0
	for _, v := range b {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	//If i don't close the channel will get deadlock error
	close(ch)
}

// takes 2 channels
// using select
func fibonacci2(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
    //if quit gets any value then print quit and break/stop
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}
//notice we didn't close the channels cause it's not nessasary here.
// cause quit handles it and there will be no overflow.

func main() {
	b := []int{3, 4, 3, 1, -5, 2}
	c := make(chan int)
	go add(b[:len(b)/2], c)
	go add(b[len(b)/2:], c)

	x, y := <-c, <-c

	//close channel will not take no more value to channnel
	close(c)

	// ok would be false if there is no value to retrieve or the channel is closed
	z, ok := <-c
	fmt.Println("check", z, ok)
	fmt.Printf("x: %v, y:%v, x+y:%v\n", x, y, x+y)

	//fibonacci
	fmt.Println("\n\nFibonacci")
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)

	for v := range ch {
		fmt.Println(v)
	}

	//fibonacci2
	fmt.Println("\n\nFibonacci only channels and select")
	newch := make(chan int)
	quit := make(chan int)

  // this gorutine is called before the fuction beacuse to set the reciever
  // if there is no reciever to get the value from the producer aka febonnaci2 func
  // then it will go to a deadlock since there is no gorutine to recieve it.
	go func() {
		for i := 0; i < 10; i++ {
			println(<-newch)
		}
		quit <- 0
	}()

	fibonacci2(newch, quit)

}
