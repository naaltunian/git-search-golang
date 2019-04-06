package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// User struct to hold JSON response from the web request
type User struct {
	Name        string `json:"name,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	URL         string `json:"url,omitempty"`
	Blog        string `json:"blog,omitempty"`
	Company     string `json:"company,omitempty"`
	Location    string `json:"location,omityempty"`
	Email       string `json:"email,omitempty"`
	Bio         string `json:"bio,omitempty"`
	PublicRepos int    `json:"public_repos,omitempty"`
	Followers   int    `json:"followers,omitempty"`
	Following   int    `json:"following,omitempty"`
}

func main() {

	url := "https://api.github.com/users/"
	var userName string

	fmt.Println("Enter a github username:")
	fmt.Scanln(&userName)

	resp, err := http.Get(url + userName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var data User

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}
