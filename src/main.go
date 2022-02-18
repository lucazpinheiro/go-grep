package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const helpMessage = `
USAGE:
	grop -f file.txt -s someword

OPTIONS:
	-h | Show instructions.  | Optional
	-f | Target file.        | Required
	-s | String to look for. | Required
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

var h = flag.Bool("h", false, "File to look for string in.")
var f = flag.String("f", "", "File to look for string in.")
var s = flag.String("s", "", "String to be searched.")

func main() {
	flag.Parse()
	if *h {
		fmt.Print(helpMessage)
		return
	}
	if *f == "" || *s == "" {
		log.Fatal("Flags '-f' and '-s' are both required argument, use -h to get help.")
	}
	file := *f
	searchedString := *s

	lines, err := readFileLineByLine(file)
	if err != nil {
		log.Fatal("Could not read file.")
	}

	for lineNumber, line := range lines {
		if strings.Contains(line, searchedString) {
			fmt.Println(lineNumber+1, line)
		}
	}
}
