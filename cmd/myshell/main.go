package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func handleType(userInputArg string) {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("type %s", userInputArg))
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s command not found\n", strings.TrimRight(userInputArg, "\n"))
	}

	fmt.Print(string(output))
}

func commandReader() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command := strings.Split(strings.TrimRight(userInput, "\n"), " ")

		switch command[0] {
		case "echo":
			fmt.Println(strings.Join(command[1:], " "))
		case "exit":
			os.Exit(0)
		case "type":
			handleType(command[1])
		default:
			fmt.Printf("%s: command not found\n", strings.TrimRight(userInput, "\n"))
		}
	}
}

func main() {
	commandReader()
}
