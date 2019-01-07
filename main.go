package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var sdtin interface{} = os.Stdin

func main() {
	host, port := parseArgs()
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		log.Panicln(err)
	}

	for {
		fmt.Print(">")
		reader := bufio.NewScanner(sdtin.(io.Reader))
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
