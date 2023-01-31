package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	fmt.Println("time.Second :", time.Second)
	time.Sleep(1 * time.Second) // sleep for 1 second
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 1048575; i++ {
		go DisplayNumber(int64(i + 1))
	}

	time.Sleep(1 * time.Second) // sleep for 10 seconds
}

func TestCreateChannel(t *testing.T) {
	var channel chan string = make(chan string)
	fmt.Println("Successfully make a channel")

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Muhammad Arfan"
		fmt.Println("Successfully send data to channel")
	}()

	name := <-channel
	fmt.Println("Successfully receive data from channel")

	fmt.Println(name)

	defer close(channel)
	fmt.Println("Successfully closing channel")

}

func TestCreateChannelBuffer(t *testing.T) {
	var channel chan string = make(chan string, 4)
	defer close(channel)

	go func() {
		channel <- "Muhammad"
		channel <- "Arfan"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Channel capacity :", cap(channel))
	fmt.Println("Channel length :", len(channel))

}

func TestCreateChannelBufferForRange(t *testing.T) {
	var channel chan string = make(chan string, 4)

	go func() {
		fullname := []string{"Muhammad", "Arfan"}
		for _, name := range fullname {
			channel <- name
			fmt.Println("Successfully set channel value :", name)
		}
		close(channel)
		fmt.Println("Channel closed")
	}()

	// go func() {
	for name := range channel {
		fmt.Println("Successfully get channel value :", name)
	}
	// }()

	fmt.Println("Channel capacity :", cap(channel))
	fmt.Println("Channel length :", len(channel))

}

func TestSelectChannel(t *testing.T) {
	var channelOne chan string = make(chan string, 4)
	var channelTwo chan string = make(chan string, 4)
	defer close(channelOne)
	defer close(channelTwo)

	go func() {
		fullname := []string{"Muhammad", "Arfan"}

		channelOne <- fullname[0]
		channelTwo <- fullname[1]
	}()

	counter := 0
	for {

		select {
		case data := <-channelOne:
			fmt.Println("data from channel one", data)
			counter++
		case data := <-channelTwo:
			fmt.Println("data from channel two", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

func TestSelectChannelWithDefault(t *testing.T) {
	var channelOne chan string = make(chan string, 4)
	var channelTwo chan string = make(chan string, 4)
	defer close(channelOne)
	defer close(channelTwo)

	go func() {
		time.Sleep(3 * time.Second)
		fullname := []string{"Muhammad", "Arfan"}

		channelOne <- fullname[0]
		channelTwo <- fullname[1]
	}()

	counter := 0
	for {

		select {
		case data := <-channelOne:
			fmt.Println("data from channel one", data)
			counter++
		case data := <-channelTwo:
			fmt.Println("data from channel two", data)
			counter++
		default:
			fmt.Println("Wating data")
		}

		if counter == 2 {
			break
		}
	}

}
