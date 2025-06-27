package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func printLine(arr []string, char int, line int) {
	str := strings.Split(arr[char-32], "\n")
	fmt.Print(str[line])
}

func main() {
	// from 32 witch is space to 126 witch is ~
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Invalid number of areguments")
		return
	}
	file, _ := os.Open("standard.txt")
	defer file.Close()
	out, _ := io.ReadAll(file)
	standart := strings.Split(string(out), "\n\n")
	input := args[0]
	inputArr := strings.Split(input, "\\n")
	for _, inp := range inputArr {
		if inp == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			// for loop check 
			for _, char := range inp {
				if char < 33 || char > 126 {
					fmt.Println("not working")
					return
				}
				if char != ' ' {
					printLine(standart, int(char), i)
				} else {
					(printLine(standart, int(' '), i+1))
				}
			}
			fmt.Println()
		}
	}
}
