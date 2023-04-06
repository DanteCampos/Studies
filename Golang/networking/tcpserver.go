package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Setup listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// Get connection and passes to handler thread
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(con)
	}
}

// Reads from connection in it's only thread
func handle(con net.Conn) {
	defer con.Close()
	dataReader := bufio.NewReader(con)
	for {
		data, err := dataReader.ReadString('\n') // Reads 100 bytes
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(data)
	}
}
