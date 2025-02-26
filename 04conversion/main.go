package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the name", name)
	fmt.Printf("The foeamt is %T\n", name)
	numString, err := strconv.ParseFloat(strings.TrimSpace(name), 64)
	if err != nil {
		fmt.Println("I am th error", err)
	} else {
		fmt.Println("Added one to the name", numString+1)
	}

}
