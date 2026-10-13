package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"slices"
	"strings"
)

var builtInCommands = []string{
	"echo",
	"pwd",
	"cd",
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
			case "pwd":
				abspath, _ := os.Getwd()
				fmt.Println(abspath)
			case "cd":
				if fields[1] == "~" {
					home, _ := os.UserHomeDir()
					os.Chdir(home)
					break
				}
				dir := path.Clean(fields[1])
				_, err := os.Stat(dir)
				if err != nil {
					fmt.Println("cd: " + fields[1] + ": No such file or directory")
					break
				}
				if err := os.Chdir(dir); err != nil {
					fmt.Println("cd: " + fields[1] + ": No such file or directory")
				}
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
