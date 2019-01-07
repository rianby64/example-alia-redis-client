package main

import (
	"fmt"
	"testing"
)

func Test_handleInput_echo_OK(t *testing.T) {

	conn := startServer(t)
	expected := "echo"

	actual, err := handleInput(conn, expected)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Error(expected, actual, "different messages")
	}

	byeServer(conn, t)
}

func Test_handleInput_bye_OK(t *testing.T) {

	conn := startServer(t)
	expected := "bye"

	_, err := handleInput(conn, expected)
	if err != nil {
		actual := fmt.Sprint(err)
		if expected != actual {
			t.Error(expected, actual, "different messages")
		}
	}

	byeServer(conn, t)
}
