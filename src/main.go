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
USAGE:
	grop [pattern] [file]
	ex: grop someword sample.txt
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
	hasPipedData := false

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() > 0 {
		hasPipedData = true
	}

	args := os.Args[1:]
	if !hasPipedData && len(args) < 2 {
		fmt.Println("Missing args, required both target file and string to search.")
		fmt.Print(helpMessage)
		os.Exit(0)
	}

	if hasPipedData && len(args) < 1 {
		log.Println("Missing pattern arg.")
		fmt.Print(helpMessage)
		os.Exit(0)
	}

	pattern := args[0]
	r, _ := regexp.Compile(pattern)

	if hasPipedData {
		// todo: search for pattern in line from pipe
		return
	}

	file := args[1]

	if _, err := os.Stat(file); err != nil && !hasPipedData {
		log.Printf("%s file not found.", file)
		os.Exit(0)
	}

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
