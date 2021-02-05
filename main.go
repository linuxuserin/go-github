package main

import (
	"fmt"
	"os"

	colour "github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

var (
	user      string
	emoji     string
	tokenfile string
)

var (
	flagUser string
	flagAuth string
	flagHelp bool
)

func main() {
	// parse flags
	flag.Parse()

	homedir, err := os.UserHomeDir()
	tokenfile = homedir + "/.go-github.token"

	if err != nil {
		colour.Red("Something went wrong")
		os.Exit(1)
	}

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage()
	}

	if flagHelp {
		printUsage()
	}

	if flagAuth != "" {
		cmdAuth()
	}
	if flagUser != "" {
		cmdUsers()
	}

}

func init() {
	flag.StringVarP(&flagUser, "user", "u", "", "Search Users")
	flag.StringVarP(&flagAuth, "auth", "x", "", "Login to Github")
	flag.BoolVarP(&flagHelp, "help", "h", false, "help message")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}
