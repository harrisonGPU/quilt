package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Message struct {
	Msg string `json:"msg"`
	Err string `json:"err"`
}

func function() {
	input := get_arg_from_caller()
	output := input + " received by the go-callee function"
	send_return_value_to_caller(output)
}

func get_arg_from_caller() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func send_return_value_to_caller(message string) {
	fmt.Print(message)
}

// func main() {
// 	function()
// }