package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func printLine(AsciiShapes []string, char int, i int) {
	str := strings.Split(AsciiShapes[char-32], "\n")
	fmt.Print(str[i])
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Invalid number of areguments... ")
		return
	}

	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error Opening File... ", err)
		return
	}
	defer file.Close()

	fileData, er := io.ReadAll(file)
	if er != nil {
		fmt.Println("Error Reading file... ", er)
		return
	}

	CleanFileData := strings.ReplaceAll(string(fileData), "\r", "")
	AsciiShapes := strings.Split(CleanFileData, "\n\n")

	input := args[0]

	// validate user input
	for _, char := range input {
		if char < 32 || char > 126 {
			fmt.Println("The string includes characters outside the ASCII range... ")
			return
		}
	}

	inputLines := strings.Split(input, "\\n")
	for _, Inpline := range inputLines {
		if Inpline == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range Inpline {
				if char != ' ' {
					printLine(AsciiShapes, int(char), i)
				} else {
					printLine(AsciiShapes, int(char), i+1)
				}
			}
			fmt.Println()
		}
	}
}
