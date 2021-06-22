package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Channels allow us to pass values between Goroutines
func main() {
	c := make(chan int, 2) // We make a channel with a buffer room of 1 so we only allow one value on the channel

	sendOnlyChannel := make(chan<- int)    // This is a directional channel which can only have ints sent to it but not pull anything from it.
	recieveOnlyChannel := make(<-chan int) // This is recieve only channel
	go func() {                            // this allows the goroutine and the main goroutine to hand off the channel data.
		c <- 42
		c <- 43
		sendOnlyChannel <- 26
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-recieveOnlyChannel)

	// You may also assign specific directed channels to a general type, but you can't convert them
	sendOnlyChannel = c

	example()
	rangeAndClosing()
	selectStatement()
	fanIn()
	contextExample()
}

func example() {
	c := make(chan int)

	// send
	go foo(c)

	//receive
	bar(c)

	fmt.Println("Main function exiting...")
}

func foo(c chan<- int) {
	c <- 26
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}

// Range and closing
func rangeAndClosing() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // closes the channel so the range doesn't get deadlocked waiting for more values.
	}()

	//receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Main function exiting...")
}

func selectStatement() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)

	fmt.Println("Ending selectStament...")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select { // looks to see which of these channels can I pull a value off of
		case v := <-e:
			fmt.Println("From even channel: ", v)
		case v := <-o:
			fmt.Println("From odd channel: ", v)
		case v, ok := <-q: // Can use the idiom ok to check if a channel is closed
			if ok {
				fmt.Println("Quit channel is now closed")
			}
			fmt.Println("From quit channel: ", v)
			return
		}
	}
}

// Fan In pattern takes values from multiple channels and pushses them into one channel.
func fanIn() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go sendFanIn(even, odd)

	go receiveFanIn(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

func sendFanIn(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receiveFanIn(e, o <-chan int, f chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			f <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			f <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(f)
}

// Take a process and split it up in many Goroutines to work all at once.
func fanOut() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanout(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Exiting fanOut...")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanout(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		// for each chunk of work launch a Goroutine to fan out and do its job
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

// Context
// This object holds request respecitive values for Go servers such that goroutines related to that request can share values like headers and cookies.
func contextExample() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("------------")

	ctx, _ = context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("------------")

}
