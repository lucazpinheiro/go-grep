package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

const helpMessage = `
USAGE:
	grop -f file.txt -s someword

OPTIONS:
	-h | Show instructions.                        | Optional
	-f | Target file.                              | Required
	-s | String to look for.                       | Required
	-c | Show total occurrences of searched string | Optional
`

func readFileLineByLine(target string) ([]string, error) {
	file, err := os.Open(target)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

func applyColor(line string, subStrLen int, positions []int) string {
	splittedLine := strings.Split(line, "")
	for _, position := range positions {
		charPosition := position
		for charPosition < subStrLen+position {
			splittedLine[charPosition] = color.YellowString(splittedLine[charPosition])
			charPosition += 1
		}
	}
	return strings.Join(splittedLine, "")
}

func findAllOccurrences(line, subString string) []int {
	lineContainsSubString := true
	subStringLen := len(subString)
	var positions []int
	for lineContainsSubString {
		subStrPosition := strings.Index(line, subString)
		if subStrPosition < 0 {
			lineContainsSubString = false
			break
		}
		positions = append(positions, subStrPosition)
		line = removeSubString(line, subStrPosition, subStringLen)
	}
	return positions
}

func removeSubString(line string, position, subStrLen int) string {
	splitedLine := strings.Split(line, "")
	for i := position; i < position+subStrLen; i++ {
		splitedLine[i] = " "
	}
	return strings.Join(splitedLine, "")
}

var h = flag.Bool("h", false, "File to look for string in.")
var c = flag.Bool("c", false, "Count substring occurrences.")
var f = flag.String("f", "", "File to look for string in.")
var s = flag.String("s", "", "String to be searched.")

func main() {
	flag.Parse()
	if *h {
		fmt.Print(helpMessage)
		return
	}
	if *f == "" || *s == "" {
		log.Fatalf("Flags '-f' and '-s' are both required argument, use -h to get help.")
	}
	file := *f
	searchedString := *s

	lines, err := readFileLineByLine(file)
	if err != nil {
		log.Fatal("Could not read file.")
	}

	count := 0
	for lineNumber, line := range lines {
		if strings.Contains(line, searchedString) {
			positions := findAllOccurrences(line, searchedString)
			if *c {
				count += len(positions)
			}
			highlightedLine := applyColor(line, len(searchedString), positions)

			fmt.Println(lineNumber+1, highlightedLine)
		}
	}
	if *c {
		fmt.Printf("%s appers %d times in file.\n", searchedString, count)
	}
}
