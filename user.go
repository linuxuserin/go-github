package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	colour "github.com/fatih/color"
)

const (
	apiURL       = "https://api.github.com"
	userEndpoint = "/users/"
)

// User Models
type User struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	Twitter           string    `json:"twitter_username"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	PrivateRepos      int       `json:"total_private_repos"`
	PrivateGists      int       `json:"private_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// GetUsers gets users details
func GetUsers(name string) User {

	if _, err := os.Stat(tokenfile); err == nil {
		tkn, err := ioutil.ReadFile(tokenfile)
		if err != nil {
			colour.Red("Something went wrong")
			os.Exit(1)
		}
		token = string(strings.Trim(string(tkn), "\n"))
	}

	req, err := http.NewRequest("GET", apiURL+userEndpoint+name, nil)
	if token != "" {
		req.Header.Add("Authorization", "token "+token)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var user User
	json.Unmarshal(body, &user)

	if (User{}) == user {

		colour.Red("Sorry the GitHub user doesn't exist")
		os.Exit(1)
		return user
	}
	return user
}
