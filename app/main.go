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
		command = strings.TrimSpace(command)

		// Exit
		if command == "exit" {
			break
		}

		// Error handling
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		// Print
		fmt.Println(command[:len(command)-1] + ": command not found")
	}
}
