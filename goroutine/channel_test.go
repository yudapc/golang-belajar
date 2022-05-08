package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// create channel, send and receive data from goroutine to channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Joko Setiawan"
	}()

	data := <-channel
	fmt.Println(data)
	close(channel)
}

// Channel as parameter

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Joko Setiawan Channel"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	close(channel)
}

// In and Out channel

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Joko Setiawan In"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)

	close(channel)
}

// Range channel
// yang dikirim secara terus menerus oleh pengirim dan tidak jelas kapan berhentinya
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("DONE")
}

// Race condition
func TestRaceConditoin(t *testing.T) {
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}
