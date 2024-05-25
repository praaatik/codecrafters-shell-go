package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("%s: command not found", strings.TrimRight(userInput, "\n"))
}
