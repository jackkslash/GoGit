package commands

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
)

func HashObject(args []string) error {
	if len(args) < 4 {
		return fmt.Errorf("not enough arguments")
	}

	if args[3] == "-w" {
		return fmt.Errorf("missing -w option")
	}

	file, err := os.ReadFile(args[3])
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	stats, err := os.Stat(args[3])
	if err != nil {
		return fmt.Errorf("failed to get file stats: %w", err)
	}

	content := string(file)
	contentAndHeader := fmt.Sprintf("blob %d\x00%s", stats.Size(), content)
	sha := (sha1.Sum([]byte(contentAndHeader)))
	hash := fmt.Sprintf("%x", sha)
	blobName := []rune(hash)
	blobPath := ".git/objects/"
	for i, v := range blobName {
		blobPath += string(v)
		if i == 1 {
			blobPath += "/"
		}
	}
	var buffer bytes.Buffer
	z := zlib.NewWriter(&buffer)
	z.Write([]byte(contentAndHeader))
	z.Close()
	os.MkdirAll(filepath.Dir(blobPath), os.ModePerm)
	f, _ := os.Create(blobPath)
	defer f.Close()
	f.Write(buffer.Bytes())
	fmt.Print(hash)
	return nil
}
