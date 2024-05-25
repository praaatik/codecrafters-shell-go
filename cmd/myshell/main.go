package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	//fmt.Printf("===")

	userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("error")
		fmt.Fprintln(os.Stderr, err)
	}
	//fmt.Printf("===")
	fmt.Printf("%s:command not found", userInput)
	//fmt.Fprintf(os.Stdout, "%s: command not found", userInput)
}
