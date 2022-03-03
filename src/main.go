package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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
	r, _ := regexp.Compile(searchedString)
	count := 0
	_, err := readFileLineByLine(file, func(line []byte) {
		positions := r.FindAllIndex(line, -1)
		occurrences := len(positions)
		if occurrences > 0 {
			if *c {
				count += occurrences
			}
			intervals := getIntervals(positions)
			highlightedLine := applyColor(line, intervals)
			fmt.Println(highlightedLine)
		}
	})
	if err != nil {
		log.Fatal("Could not read file.")
	}
	if *c {
		fmt.Printf("\nfound %d matchs for patter %s in file.\n", count, searchedString)
	}
}
