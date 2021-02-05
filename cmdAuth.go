package main

import (
	"io/ioutil"
	"os"

	colour "github.com/fatih/color"
)

var (
	token string
)

func cmdAuth() {
	token = flagAuth

	if token == "" {
		colour.Red("Error: token empty")
		os.Exit(0)
	}
	if len(token) != 40 {
		colour.Red("Error: Invalid token")
		os.Exit(0)
	}

	if flagUser == "" && flagAuth != "" {
		if _, err := os.Stat(tokenfile); os.IsNotExist(err) {
			err = ioutil.WriteFile(tokenfile, []byte(token), 0644)
			if err != nil {
				colour.Red("Something went wrong")
				os.Exit(1)
			}
		}
	}
}
