package main

import (
	"fmt"
	"time"
)

//single channel

func CreateWorker(id int) chan<- int {
	c := make(chan int)

	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i)
		go CreateWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Minute)
}

//Buffer channel
func BufferWorker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}

func BufferCreateWorker(id int, c chan int) chan<- int {
	//c := make(chan int)

	go BufferWorker(id, c)
	return c
}
func bufferedChannel() {
	c := make(chan int, 3)
	go BufferCreateWorker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

//channelClose
func CloseWorker(id int, c <-chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}
func ChannelClose() {
	c := make(chan int)
	go CloseWorker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	close(c)
	//time.Sleep(time.Millisecond)
}
func main() {
	chanDemo()
	bufferedChannel()
	ChannelClose()
	time.Sleep(time.Millisecond)
}
