package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

// building a TCP echo server listening on a port to echo back to client
// Test: echo Hello world | nc localhost 9090 using netcat command, send request to server
func main() {
	// go run main.go 9090
	// arg[0] file name
	// arg[1] port
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go <port>")
		os.Exit(1)
	}

	port := fmt.Sprintf(":%s", os.Args[1])

	listener, err := net.Listen("tcp", port) // build a TCP listener on the inout port
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
	}
	defer listener.Close()

	fmt.Println("listening on:", listener.Addr())

	// accept the connections on the listener
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue // keep listening
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn) // efficiently handle the conenction content in buffers
	for {
		bytes, err := reader.ReadBytes(byte('\n')) // read from buffer for each line
		if err != nil {
			if err != io.EOF { // EOF: client gracefully ended connection or file ended
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s\n", bytes) // equivalent: fmt.Println("request:", string(bytes))

		line := fmt.Sprintf("Echo: %s", bytes)

		fmt.Printf("response: %s\n", line)

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("failed to write data, err:", err)
			return
		}
	}

}
