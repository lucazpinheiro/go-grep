package main

import (
	"flag"
	"fmt"
	"log"
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
	// file, err := os.Open(target)
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)

	var text []string
	x := []byte("'Cês vão escutar falar dos cara toda hora")
	fmt.Println("len x", len(x))
	fmt.Println("len x string", len(string(x)))
	callback(x)
	// for scanner.Scan() {
	// 	callback(scanner.Bytes())
	// }

	return text, nil
}

func applyColor(line string, positions [][]int) string {
	fmt.Println(len(line))
	splittedLine := strings.Split(line, "")
	fmt.Println(splittedLine)
	for _, position := range positions {
		currentPosition := position[0]
		finalPosition := position[1]

		for currentPosition < finalPosition {
			if currentPosition >= len(splittedLine) {
				splittedLine[currentPosition-1] = color.YellowString(splittedLine[currentPosition-1])
				currentPosition += 1
				continue
			}
			splittedLine[currentPosition] = color.YellowString(splittedLine[currentPosition])
			if (currentPosition + 1) == len(splittedLine) {
				currentPosition = finalPosition
				continue
			}
			currentPosition += 1
		}

	}

	return strings.Join(splittedLine, "")
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

	_, err := readFileLineByLine(file, func(line []byte) {
		positions := r.FindAllIndex(line, -1)
		if len(positions) > 0 {
			highlightedLine := applyColor(string(line), positions)
			fmt.Println(highlightedLine)
		}
	})
	if err != nil {
		log.Fatal("Could not read file.")
	}
}
