package main

import (
	"fmt"
	"os"

	"github.com/jackkslash/GoGit/commands"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: gogit <command> [<args>...]\n")
		os.Exit(1)
	}

	var err error
	switch command := os.Args[1]; command {
	case "init":
		err = commands.Init(os.Args)
	case "cat-file":
		err = commands.CatFile(os.Args)
	case "help":
		err = commands.Help(os.Args)
	default:
		err = fmt.Errorf("unknown command %s", command)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
