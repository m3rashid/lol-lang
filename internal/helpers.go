package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const FILE_EXTENSION = "lol"

func GetFileContents(filePath string) ([]string, error) {
	if filePath == "" {
		return nil, errors.New("error: Please provide file path \nUsage: lol <file-path>")
	}

	fileExtension := strings.Split(filePath, ".")
	if fileExtension[len(fileExtension)-1] != FILE_EXTENSION {
		return nil, fmt.Errorf("error: Invalid file extension. Please provide a .%s file", FILE_EXTENSION)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Trim(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
