package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		// User input
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		}

		if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
		} else {
			fmt.Println(command + ": command not found")
		}

		if strings.HasPrefix(command, "type ") {
			command = strings.Trim(command, "type")
			command = strings.TrimSpace(command)
			switch {
			case command == "exit" || command == "echo" || command == "type":
				fmt.Println(command, "is a shell builtin")
			default:
				fmt.Println(command, ": not found")
			}
		}
	}
}
