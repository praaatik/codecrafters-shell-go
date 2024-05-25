package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandReader() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Printf("%s: command not found\n", strings.TrimRight(userInput, "\n"))
	}

}

func main() {
	commandReader()
}
