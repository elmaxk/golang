package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Defines an io.Reader to read from stdin
type MaxReader struct{}

// Reads data from stdin
func (maxReader *MaxReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// Writer defines an io.Writer to write to Stdout.
type MaxWriter struct{}

// Writes data to Stdout
func (maxWriter *MaxWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

// func Copy(dst io.Writer, src io.Reader) (written int64, error)

func main() {
	// Instantiate reader and writer
	var (
		reader MaxReader
		writer MaxWriter
	)

	// Create buffer to hold input/output
	input := make([]byte, 4096)

	// Use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// Use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}

	fmt.Printf("Wrote %d bytes to stdout\n", s)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
