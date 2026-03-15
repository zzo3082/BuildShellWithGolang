package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	for {
		fmt.Print("$ ")
		comment, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		userinput := strings.TrimSpace(comment)
		// exit command to break the loop
		if userinput == "exit" {
			break
		}

		// echo command to print the input back to the user
		if strings.HasPrefix(userinput, "echo") {
			echoText := strings.TrimPrefix(userinput, "echo ")
			fmt.Println(echoText)
		} else {
			fmt.Println(userinput + ": command not found")
		}
	}
}
