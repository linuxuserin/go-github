package main

import (
	"fmt"
	"strings"

	colour "github.com/fatih/color"
)

func cmdUsers() {
	users := strings.Split(flagUser, ",")
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
		colour.Green(`Website:	%s`, result.Blog)
		colour.Green(`Twitter:		%s`, result.Twitter)
		colour.Green(`Email:		%s`, result.Email)
		colour.Blue(`Location:	%s`, result.Location)
		colour.Blue(`Company:	%s`, result.Company)
		colour.Blue(`Repositories:	%d(s)`, result.PublicRepos+result.PrivateRepos)
		colour.Blue(`Gists:	%d(s)`, result.PublicGists+result.PrivateGists)
		colour.Green(`Followers:		%d(s)`, result.Followers)
		colour.Green(`Following:		%d(s)`, result.Following)
		fmt.Println("")
	}
}
