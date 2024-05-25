package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func commandReader() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command := strings.Split(strings.TrimRight(userInput, "\n"), " ")

		switch command[0] {
		case "echo":
			//handleEcho(command[1:])
			fmt.Println(strings.Join(command[1:], " "))
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("%s: command not found\n", strings.TrimRight(userInput, "\n"))
		}
	}

}

func main() {
	commandReader()
}
