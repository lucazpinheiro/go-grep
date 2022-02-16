package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

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

var file = flag.String("f", "", "File to look for string in.")
var str = flag.String("s", "", "String to be searched.")

func main() {
	flag.Parse()
	if *file == "" {
		panic("'-f' flag is a required argument")
	}
	file := *file

	if *str == "" {
		panic("'-s' flag is a required argument")
	}
	searchedString := *str

	lines, err := readLineByLine(file)
	if err != nil {
		panic("deu ruim")
	}

	for lineNumber, line := range lines {
		if strings.Contains(line, searchedString) {
			fmt.Println(lineNumber+1, line)
		}
	}
	// content, size, err := getFileContent(testFile)
	// if err != nil {
	// 	panic("deu ruim")
	// }

	// for i, line := range content {
	// 	fmt.Println(i, string(line))
	// }
	// fmt.Println(content)
	// fmt.Println("size: ", size)

}
