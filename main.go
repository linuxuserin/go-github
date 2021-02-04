package main

import (
	"fmt"
	"os"
	"strings"

	colour "github.com/fatih/color"
	flag "github.com/ogier/pflag"
)

var (
	user  string
	emoji string
)

func main() {
	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage()
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	fmt.Println("")
	for _, u := range users {
		result := GetUsers(u)
		colour.Cyan(`Username:	%s`, result.Login)
		colour.Cyan(`ID:		%d`, result.ID)
		colour.Cyan(`Created on:	%s`, result.CreatedAt)
		colour.Cyan(`Updated on:	%s`, result.UpdatedAt)
		colour.Blue(`Name:		%s`, result.Name)
		colour.HiMagenta(`Bio:		%s`, result.Bio)
		colour.Green(`Website:		%s`, result.Blog)
		colour.Green(`Twitter:		%s`, result.Twitter)
		colour.Green(`Email:		%s`, result.Email)
		colour.Blue(`Location:	%s`, result.Location)
		colour.Blue(`Company:	%s`, result.Company)
		colour.Blue(`Repositories:	%d(s)`, result.PublicRepos)
		colour.Blue(`Gists:	%d(s)`, result.PublicGists)
		colour.Green(`Followers:		%d(s)`, result.Followers)
		colour.Green(`Following:		%d(s)`, result.Following)
		fmt.Println("")
	}

}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
