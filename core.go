package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func convertPxToRem(config config, line string) string {

	// split by space
	bySpace := strings.Split(line, " ")

	newLine := ""
	for _, content := range bySpace {
		if strings.Contains(content, "px") {
			i := strings.Index(content, "px")
			numberString := content[:i]
			number, err := strconv.ParseFloat(numberString, 64)

			if err != nil {
				continue
			}
			newNumber := config.conversionFactor * number

			newString := fmt.Sprintf("%.1frem", newNumber)

			endString := content[i+2:]

			if len(endString) > 0 {
				newLine += newString + endString
			} else {
				newLine += newString + " "
			}
		} else {
			newLine += content + " "
		}

	}

	return newLine
}

func parseContents(config config, contents string) string {
	newContents := ""

	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		// if line contains px, split it for every px

		if strings.Contains(line, "px") &&
			strings.Contains(line, config.doNotInclude) {
			newContents += convertPxToRem(config, line)
		} else {
			newContents += line
		}

		newContents += "\n"
	}

	return newContents
}

type config struct {
	conversionFactor float64
	doNotInclude     string
}

func charmInterface(config config) error {
	root := "." // Starting directory
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, "css") {
			contents, _ := os.ReadFile(path)

			newContents := parseContents(config, string(contents))

			os.WriteFile(path, []byte(newContents), 0755)
		}

		return nil
	})

	return err
}
