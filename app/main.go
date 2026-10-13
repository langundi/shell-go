package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var builtInCommands = []string{
	"echo",
	"type",
	"exit",
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		fields := strings.Fields(command)

		if len(fields[0]) == 0 {
			os.Exit(1)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		if slices.Contains(builtInCommands, fields[0]) {
			switch fields[0] {
			case "exit":
				os.Exit(0)
			case "echo":
				fmt.Println(command[5:])
			case "type":
				switch fields[1] {
				case "exit", "echo", "type":
					fmt.Println(fields[1], "is a shell builtin")
				default:
					path, err := exec.LookPath(fields[1])
					if err != nil {
						fmt.Println(fields[1] + ": not found")
						break
					}
					fmt.Println(fields[1], "is", path)
				}
			}
		} else if _, err := exec.LookPath(fields[0]); err == nil {
			cmd := exec.Command(fields[0], fields[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			fmt.Println(command + ": command not found")
		}
	}
}
