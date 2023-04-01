package main

import (
	"flag"
	"os"
	"bbg-cli/cmd"
)

func main() {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	themeURL := initCmd.String("theme", "https://github.com/bbg-contributors/default_theme_src/raw/main/index.html", "the URL of BBG theme")

	if len(os.Args) < 2 {
		cmd.Help()
	}

	args := os.Args[2:]

	switch os.Args[1] {

	case "init":
		initCmd.Parse(args)
		cmd.Init(*themeURL)
	
	default:
		cmd.Help()
	}
}
