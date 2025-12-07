package inOut

import (
	"os"
)

// Read reads content from a file inside the inOut directory
func Read(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Write writes content to a file inside the inOut directory
func Write(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}
