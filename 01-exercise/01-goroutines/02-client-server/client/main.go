package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8000")
	defer con.Close()
	if err != nil {
		log.Fatal(err)
	}

	mustCopy(os.Stdout, con)
}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
