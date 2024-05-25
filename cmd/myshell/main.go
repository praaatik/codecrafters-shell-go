package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func handleType(userInputArg string) {
	// hard coding the values, for now
	if userInputArg == "echo" || userInputArg == "exit" || userInputArg == "type" {
		fmt.Printf("%s is a shell builtin\n", userInputArg)
		return
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("type -P %s", userInputArg))
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s: not found\n", userInputArg)
		return
	}

	fmt.Printf("%s is %s", userInputArg, string(output))
	return
}

func handleExecution(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", cmd)
	}
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
			handleExecution(command[0], command[1:])
		}
	}
}

func main() {
	commandReader()
}
