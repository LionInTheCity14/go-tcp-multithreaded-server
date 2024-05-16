package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // schedule its close before function finish

	// Read the request from the client
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal("Error reading:", err)
	}

	fmt.Println("processing the request")
	time.Sleep(1 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n Successfully received response from server!! \r\n"))
}

func main() {
	port := 5000
	// listen for incomming connections on port 5000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	//  schedule a listener.Close() call to be executed just before main function returns
	defer listener.Close()
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	for {
		fmt.Println()
		fmt.Println("Server is listening on port:", port)

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection:", err)
		}

		handleConnection(conn)
	}
}
