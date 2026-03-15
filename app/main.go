package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

		// 1. 將輸入拆分成 指令(cmdName) 和 參數(args)
		// strings.Fields 會自動幫你處理多個空格的問題，比 strings.Split(" ") 更安全
		parts := strings.Fields(userinput)
		cmdName := parts[0]
		args := parts[1:] // 取出第一個元素之後的所有內容當作參數

		// exit command : break the loop
		if cmdName == "exit" {
			break
		}

		// echo command : print the input back to the user
		if cmdName == "echo" {
			fmt.Println(strings.Join(args, " "))
		}

		// type command : check fllowing string is valid command or not
		if cmdName == "type" {
			if len(args) == 0 {
				continue
			}
			typeTarget := args[0]
			if slices.Contains(validCommands, typeTarget) {
				fmt.Println(typeTarget + " is a shell builtin")
			} else if path, err := exec.LookPath(typeTarget); err == nil {
				fmt.Println(typeTarget + " is " + path)
			} else {
				fmt.Println(typeTarget + ": not found")
			}
			continue
		}

		// 3. 處理外部程式 (External Programs)
		// 如果上面的內建指令都沒命中，程式就會走到這裡

		// 檢查指令是否存在於 PATH 中
		_, err = exec.LookPath(cmdName)
		if err != nil {
			// 真的找不到，才印出 command not found
			fmt.Println(userinput + ": command not found")
			continue
		}

		// 指令存在！準備執行它
		// exec.Command(指令名稱, 參數1, 參數2...)
		cmd := exec.Command(cmdName, args...)

		// 將這個外部程式的輸出/輸入/錯誤，直接綁定到我們目前的 Shell 畫面上
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run() 會啟動程式並等待它執行結束
		err = cmd.Run()
		if err != nil {
			// 如果執行過程中發生錯誤（例如沒有權限等）
			fmt.Println("Error executing program:", err)
		}

	}
}
