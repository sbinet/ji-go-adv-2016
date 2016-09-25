// This program extends part 1.
//
// It makes a connection the host and port specified by the -dial flag, reads
// lines from standard input and writes JSON-encoded messages to the network
// connection.
//
// You can test this program by installing and running the dump program:
// 	$ go get github.com/sbinet/whispering-gophers/util/dump
// 	$ dump -listen=localhost:8000
// And in another terminal session, run this program:
// 	$ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
)

var dialAddr = flag.String("dial", "localhost:8000", "host:port to dial")

type Message struct {
	Body string
}

func main() {
	// TODO: Parse the flags.

	// TODO: Open a new connection using the value of the "dial" flag.
	// TODO: Don't forget to check the error.
	// TODO: Don't forget to close the connection (using a defer statement.)

	s := bufio.NewScanner(os.Stdin)
	// TODO: Create a json.Encoder writing into the connection you created before.
	for s.Scan() {
		m := Message{Body: s.Text()}
		err := e.Encode(m)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
