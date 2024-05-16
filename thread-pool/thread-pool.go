package main

import (
	"fmt"
	"net"
	"sync"
)

const (
	maxWorkers = 10 // Maximum number of workers in the pool
)

func main() {
	// Initialize a wait group to synchronize the termination of all workers
	var wg sync.WaitGroup

	// Create a channel to receive incoming connections
	connections := make(chan net.Conn)

	// Start the worker pool
	for i := 0; i < maxWorkers; i++ {
		go worker(connections, &wg)
	}

	// Listen for incoming connections
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 5000...")

	// Accept connections in a loop and send them to the worker pool
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		connections <- conn
	}

	// Close the connections channel when the main loop exits
	close(connections)

	// Wait for all workers to finish
	wg.Wait()
}

// worker function that handles connections
func worker(connections <-chan net.Conn, wg *sync.WaitGroup) {
	// Increment the wait group counter when starting a new worker
	wg.Add(1)

	// Decrement the wait group counter when the worker finishes
	defer wg.Done()

	for conn := range connections {
		// Handle the connection here
		fmt.Printf("Handling connection from %s\n", conn.RemoteAddr())

		// For demonstration purposes, just close the connection
		conn.Close()
	}
}
