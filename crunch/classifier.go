package crunch

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Classifier struct {
	mappings map[string]string
}

func NewClassifier(path string) (*Classifier, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()

	scanner := bufio.NewScanner(file)

	isAwaitingClassificationName := true
	var currentClassificationName string
	mappings := make(map[string]string)

	for scanner.Scan() {
		currentLine := strings.TrimSpace(scanner.Text()) // Indentation doesn't matter

		if currentLine == "" {
			isAwaitingClassificationName = true
			continue
		}

		if isAwaitingClassificationName {
			currentClassificationName = currentLine
			isAwaitingClassificationName = false
			continue
		} else {
			entry := currentLine
			mappings[entry] = currentClassificationName
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Classifier{mappings: mappings}, nil
}

func (c *Classifier) Classify(input string) (string, error) {
	for pattern, classification := range c.mappings {
		matched, err := regexp.MatchString(pattern, input)
		if err != nil {
			log.Fatal(err)
		}

		if matched {
			return classification, nil
		}
	}

	return "", fmt.Errorf("could not classify: %s", input)
}
