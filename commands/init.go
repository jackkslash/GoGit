package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Init(args []string) error {
	path := "."
	if len(args) == 3 {
		path = args[2]
	}

	if path != "." && path != "test" {
		return fmt.Errorf("init can only be used with '.' or 'test' as the directory")
	}

	if isGitRepo(path) {
		return fmt.Errorf("reinitialization of existing git repository is not allowed")
	}

	if err := initRepo(path); err != nil {
		return fmt.Errorf("error initializing repository: %w", err)
	}

	fmt.Printf("Initialized git directory in %s\n", path)
	return nil
}

func initRepo(path string) error {
	dirs := []string{".git", ".git/objects", ".git/refs"}
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(path, dir), 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", filepath.Join(path, dir), err)
		}
	}

	headFilePath := filepath.Join(path, ".git", "HEAD")
	if err := os.WriteFile(headFilePath, []byte("ref: refs/heads/main\n"), 0644); err != nil {
		return fmt.Errorf("error writing file %s: %w", headFilePath, err)
	}

	return nil
}

func isGitRepo(path string) bool {
	gitDir := filepath.Join(path, ".git")
	_, err := os.Stat(gitDir)
	return err == nil
}
