package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

const helpMessage = `
USAGE: grop [pattern] [file]
	ex: grop func main.go
		grop '[A-Z][a-z]' sample.txt
`

func readFileLineByLine(target string, callback func([]byte)) ([]string, error) {
	file, err := os.Open(target)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		callback(scanner.Bytes())
	}

	return text, nil
}

func getIntervals(positions [][]int) []int {
	var intervals []int
	for _, position := range positions {
		initialPosition := position[0]
		finalPosition := position[1]
		for initialPosition < finalPosition {
			intervals = append(intervals, initialPosition)
			initialPosition += 1
		}
	}
	return intervals
}

func intervalContainsPosition(interval []int, position int) bool {
	for _, elem := range interval {
		if position == elem {
			return true
		}
	}
	return false
}

func applyColor(line []byte, intervals []int) string {
	var stringifiedLine []string

	for charPosition, char := range string(line) {
		if intervalContainsPosition(intervals, charPosition) {
			stringifiedLine = append(stringifiedLine, color.YellowString(string(char)))
			continue
		}
		stringifiedLine = append(stringifiedLine, string(char))
	}

	return strings.Join(stringifiedLine, "")
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Missing args, both pattern and target file are required.")
		fmt.Print(helpMessage)
		os.Exit(0)
	}

	if len(args) > 2 {
		fmt.Println("Too much args were guiven, only pattern and target file are required.")
		fmt.Print(helpMessage)
		os.Exit(0)
	}

	pattern := args[0]
	file := args[1]

	if _, err := os.Stat(file); err != nil {
		log.Printf("%s file not found.", file)
		os.Exit(0)
	}

	r, _ := regexp.Compile(pattern)

	_, readingLineErr := readFileLineByLine(file, func(line []byte) {
		positions := r.FindAllIndex(line, -1)
		occurrences := len(positions)
		if occurrences > 0 {
			intervals := getIntervals(positions)
			highlightedLine := applyColor(line, intervals)
			fmt.Println(highlightedLine)
		}
	})
	if readingLineErr != nil {
		log.Fatal("Could not read file.")
	}
}
