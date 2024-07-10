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
			foundBeginning := false
			i := strings.Index(content, "px")
			begString := content[:i]
			numberStringStart := strings.Index(begString, "(")
			if numberStringStart == -1 {
				numberStringStart = 0
			} else {
				numberStringStart += 1
				foundBeginning = true
			}
			numberString := begString[numberStringStart:]

			number, err := strconv.ParseFloat(numberString, 64)

			if err != nil {
				newLine += content + " "
				continue
			}
			newNumber := roundFloat(number/config.conversionFactor, uint(config.precision))

			newString := fmt.Sprintf("%.*frem", config.precision, newNumber)

			endString := content[i+2:]

			if foundBeginning {
				newLine += begString[:len(begString)-len(numberString)]
			}
			if len(endString) > 0 {
				newLine += newString + endString
			} else {
				newLine += newString + " "
			}
		} else {
			newLine += content + " "
		}

	}

	newLineWithSpaces := ""
	for i, c := range newLine {
		if c == ',' && i+1 < len(newLine) && newLine[i+1] != ' ' {
			newLineWithSpaces += ", "
		} else {
			newLineWithSpaces += fmt.Sprintf("%c", c)
		}

	}

	return strings.TrimSpace(newLineWithSpaces)
}

func checkInclusion(config config, line string) string {
	fail := false
	if config.doNotInclude != "" {
		vals := strings.Split(config.doNotInclude, ",")
		for _, str := range vals {
			if strings.Contains(line, strings.TrimSpace(str)) {
				fail = true
			}
		}
	}

	if strings.Contains(line, "px") && !fail {
		newLine := convertPxToRem(config, line)
		first := greenFill.Render(strings.TrimSpace(newLine))
		second := primaryFill.Render(strings.TrimSpace(line))
		fmt.Println(second + " -> " + first)

		return newLine
	} else {
		return line
	}
}

func parseContents(config config, contents string) string {
	newContents := ""

	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		// if line contains px, split it for every px

		newContents += checkInclusion(config, line)

		newContents += "\n"
	}

	return newContents
}

type config struct {
	conversionFactor float64
	doNotInclude     string
	fileExtension    string
	precision        int
}

func printConfig(config config) {
	fmt.Println(config.conversionFactor)
	fmt.Println(config.doNotInclude)
	fmt.Println(config.fileExtension)
	fmt.Println(config.precision)
}

func charmInterface(config config) error {

	printConfig(config)

	root := "." // Starting directory
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, "node_modules") {
			return nil
		}

		for _, fileExtensionList := range strings.Split(config.fileExtension, ",") {
			if strings.Contains(path, strings.TrimSpace(fileExtensionList)) {
				contents, _ := os.ReadFile(path)

				fmt.Println(fileFill.Render(path + ": "))
				newContents := parseContents(config, string(contents))

				os.WriteFile(path, []byte(newContents), 0755)
				fmt.Println()
			}

		}

		return nil
	})

	return err
}
