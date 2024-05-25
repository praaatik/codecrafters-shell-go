package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isExit(inputString string) bool {

	if strings.Split(inputString, " ")[0] == "exit" {
		return true
	}
	return false
}

func commandReader() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if isExit(userInput) {
			break
		}
		fmt.Printf("%s: command not found\n", strings.TrimRight(userInput, "\n"))
	}

}

func main() {
	commandReader()
}
