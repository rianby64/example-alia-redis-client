package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	// connect to this socket
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		log.Panicln(err)
	}

	for {
		fmt.Print(">")
		// read in input from stdinh
		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		message := reader.Text()

		_, err := conn.Write([]byte(fmt.Sprintln(message)))
		if err != nil {
			log.Panicln(err)
		}

		// listen for reply
		responser := bufio.NewScanner(conn)
		if responser.Scan() {
			response := responser.Text()
			fmt.Println(response)
			if response == "bye" {
				break
			}
		} else {
			break
		}
	}
}
