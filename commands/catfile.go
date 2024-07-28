package commands

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"strings"
)

func CatFile(args []string) error {
	if args[2] != "-p" {
		fmt.Printf("Refer the help for catfile command\n")
		return nil
	}

	sha := args[3]
	path := fmt.Sprintf(".git/objects/%v/%v", sha[:2], sha[2:])
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening file %s: %w", path, err)
	}
	defer file.Close()
	reader, err := zlib.NewReader(io.Reader(file))
	if err != nil {
		return fmt.Errorf("error creating zlib reader: %w", err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	parts := strings.Split(string(data), "\x00")
	if len(parts) < 2 {
		return fmt.Errorf("unxpected data format")
	}
	fmt.Print(parts[1])
	return nil
}
