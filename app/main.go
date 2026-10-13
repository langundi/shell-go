package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		} else if strings.HasPrefix(command, "type ") {
			command = strings.TrimPrefix(command, "type ")
			command = strings.TrimSpace(command)

			switch command {
			case "exit", "echo", "type":
				fmt.Println(command, "is a shell builtin")
			default:
				path, err := exec.LookPath(command)
				if err != nil {
					fmt.Println(command + ": not found")
					break
				}
				fmt.Println(command, "is", path)
			}
		} else {
			path, err := exec.LookPath(command)
			if err != nil {
				fmt.Println(command + ": not found")
				break
			}
			fmt.Println(command, "is", path)
			fmt.Println(command + ": command not found")
		}
	}
}
