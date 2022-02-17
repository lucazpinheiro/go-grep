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

func getFileContent(target string) ([]byte, int, error) {
	file, err := os.Open(target)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return nil, 0, err
	}

	linesInFile := int(stat.Size())

	// read the file
	contentBytes := make([]byte, linesInFile)
	_, err = file.Read(contentBytes)
	if err != nil {
		return nil, 0, err
	}

	return contentBytes, linesInFile, nil
}

func readLineByLine(target string) ([]string, error) {
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

func parseArgs() (string, string) {
	args := os.Args[1:]
	if len(args) < 2 {
		panic("missing arguments")
	}
	return args[0], args[1]
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

	lines, err := readLineByLine(file)
	if err != nil {
		log.Fatal("deu ruim")
	}

	for lineNumber, line := range lines {
		if strings.Contains(line, searchedString) {
			fmt.Println(lineNumber+1, line)
		}
	}
}
