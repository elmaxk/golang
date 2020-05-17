package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

// Handler that echoes received daa
func echo(conn net.Conn) {
	defer conn.Close()

	/*
		// Create a buffer to store received data.
		b := make([]byte, 512)

		for {
			// Receive data via conn.Read into a buffer.
			size, err := conn.Read(b[0:])
			if err == io.EOF {
				log.Println("Client disconnected")
				break
			}
			if err != nil {
				log.Println("Unexpected error")
				break
			}
			log.Printf("Received %d bytes: %s\n\n", size, string(b))

			// Send data via conn.Write
			log.Println("Writing data")
			if _, err := conn.Write(b[0:size]); err != nil {
				log.Fatalln("Unable to write data")
	*/

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s", len(s), s)

	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

func echo2(conn net.Conn) {
	defer conn.Close()
	// Copy data from io.Reader to io.Writer via io.Copy().
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data.")
	}
}

func main() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}
	for {
		// Wait for a connection. Create net.Conn on connection establishment
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency.
		go echo(conn)

	}
}
