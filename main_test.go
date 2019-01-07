package main

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func startServer(t *testing.T) *net.Conn {

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		t.Error(err)
	}
	go (func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Fatal(err)
			}
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				text := scanner.Text()
				conn.Write([]byte(fmt.Sprintln(text)))
				if text == "bye" {
					return
				}
			}
		}
	})()

	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		t.Fatal(err)
	}
	return &conn
}

func byeServer(conn *net.Conn, t *testing.T) {
	if _, err := (*conn).Write([]byte(fmt.Sprintln("bye"))); err != nil {
		t.Fatal(err)
	}
}

func Test_Echo_Server_Mock_OK(t *testing.T) {

	conn := *startServer(t)
	expected := fmt.Sprintln("echo")

	if _, err := conn.Write([]byte(fmt.Sprintln(expected))); err != nil {
		t.Fatal(err)
	}

	reader := bufio.NewReader(conn)
	actual, err := reader.ReadString('\n')

	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatal(expected, actual, "different messages")
	}

	byeServer(&conn, t)
}
