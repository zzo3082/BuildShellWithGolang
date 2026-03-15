package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print
var validCommands = []string{"echo", "exit", "type"}

func main() {
	for {
		fmt.Print("$ ")
		comment, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		userinput := strings.TrimSpace(comment)
		// 先判斷指令是不是valid的
		command := strings.Split(userinput, " ")[0]
		if !slices.Contains(validCommands, command) {
			fmt.Println(userinput + ": command not found")
			continue
		}

		// exit command : break the loop
		if userinput == "exit" {
			break
		}

		// echo command : print the input back to the user
		if strings.HasPrefix(userinput, "echo") {
			echoText := strings.TrimPrefix(userinput, "echo ")
			fmt.Println(echoText)
		}

		// type command : check fllowing string is valid command or not
		if strings.HasPrefix(userinput, "type") {
			typeText := strings.TrimPrefix(userinput, "type ")
			if slices.Contains(validCommands, typeText) {
				fmt.Println(userinput + " is a shell builtin")
			} else {
				fmt.Println(userinput + ": not found")
			}
		}
	}
}
