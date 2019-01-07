package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var sdtin = os.Stdin

func main() {
	host, port := parseArgs()

	// connect to this socket
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		log.Panicln(err)
	}

	for {
		fmt.Print(">")
		// read in input from stdinh
		reader := bufio.NewScanner(sdtin)
		reader.Scan()
		message := reader.Text()

		response, err := handleInput(&conn, message)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(response)
	}
}
