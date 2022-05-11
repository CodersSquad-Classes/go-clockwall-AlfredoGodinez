// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	//"os"
	//"strings"
	"flag"
)

//This declares an integer flag, -n, stored in the pointer nFlag, with type *int:
var port = flag.Int("port", 8000, "Server port")

func handleConn(c net.Conn) {
	//flag.Parse()		//to parse the command line into the defined flags.
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse() //to parse the command line into the defined flags.
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}
