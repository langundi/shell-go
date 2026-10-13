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
				case "exit", "echo", "type", "pwd":
					fmt.Println(fields[1], "is a shell builtin")
				default:
					result, err := checkPath(fields[1])
					if err != nil {
						fmt.Println(result)
						break
					}
					fmt.Println(result)
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

func checkPath(field string) (string, error) {
	path, err := exec.LookPath(field)
	if err != nil {
		return field + ": not found", err
	}
	return fmt.Sprintf("%s is %s", field, path), nil
}
