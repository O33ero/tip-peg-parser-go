package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tip-peg-parser-go/pkg"
)

func main() {
	code := "x = 1;\n" +
		"y = 2;\n" +
		"z = x + y;\n" +
		"output z;"

	rootStatement := pkg.Parse(strings.Split(code, "\n"))
	fmt.Printf("%s", rootStatement)
}

func openFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, err
}

func getFileLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	textBox := scanner.Text()
	lines := strings.Split(textBox, ",")
	return lines
}
