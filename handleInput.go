package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func handleInput(conn *net.Conn, message string) (string, error) {
	_, err := (*conn).Write([]byte(fmt.Sprintln(message)))
	if err != nil {
		return "", err
	}

	// listen for reply
	var response string
	responser := bufio.NewScanner(*conn)
	if responser.Scan() {
		response = responser.Text()
		if response == "bye" {
			return "", errors.New("bye")
		}
	}
	return response, nil
}
