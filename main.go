package main

import (
	"bufio"
	"fmt"
	"os"
	// "encoding/json"
)

func main() {

	// url := "https://api.github.com/users/"
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a github username:")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
