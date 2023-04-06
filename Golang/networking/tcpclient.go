package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0

// Create threads and wait for them to finish
func main() {
	for i := 0; i < 200; i++ {
		wg.Add(1)
		counter++
		go sender(i)
	}
	go monitor()
	wg.Wait()
}

// Prints how many threads are executing
func monitor() {
	for {
		fmt.Printf("%v threads executing\n", counter)
		time.Sleep(time.Millisecond)
	}
}

// Dials a connection and sends messages
func sender(id int) {
	defer wg.Done()

	con, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()

	for i := 0; i < 100; i++ {
		msg := []byte(strconv.Itoa(id) + " - " + time.Now().String() + "\n")
		con.Write(msg)
		time.Sleep(20 * time.Millisecond)
	}
	msg := []byte(strconv.Itoa(id) + " ended connection\n")
	con.Write(msg)

	counter--
}
