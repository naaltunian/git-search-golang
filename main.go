package main

import (
	"fmt"
	// "encoding/json"
)

func main() {

	// url := "https://api.github.com/users/"
	var userName string

	fmt.Println("Enter a github username:")
	fmt.Scanln(&userName)
	fmt.Printf("Username: %s\n", userName)
}
