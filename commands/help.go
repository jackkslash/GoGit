package commands

import (
	"fmt"
)

func Help(args []string) error {

	switch command := args[1]; command {
	case "init":
		fmt.Printf("Initialize a git repository in the current directory or in a specified directory\n")
		fmt.Printf("Usage: gogit init [<directory>]\n")
	case "catfile":
		fmt.Printf("Print the contents of a file in the repository\n")
		fmt.Printf("Usage: gogit catfile <file>\n")
	case "help":
		helpText := `usage: gogit init [<directory>]

Create an empty Git repository or reinitialize an existing one.

Arguments:
   <directory>   The directory to create the repository in. 
                 Defaults to current directory using "." if not specified.`

		fmt.Println(helpText)
	default:
		return fmt.Errorf("unknown command %s", command)
	}

	return nil
}
