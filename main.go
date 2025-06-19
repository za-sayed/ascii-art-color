package main

import (
	"ascii-art/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // All arguments except the program name
	var style string

	// Ensure that enough arguments are passed
	if len(args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=red kit \"a king kitten have kit\"")
		return
	} else if len(args) == 2 {
		style = "standard"
	} else if len(args) == 3 && strings.ToLower(os.Args[3]) == "shadow" || strings.ToLower(os.Args[3]) == "standard" || strings.ToLower(os.Args[3]) == "thinkertoy" {
		style = strings.ToLower(os.Args[3]) // Check if there are 5 arguments before accessing os.Args[4]
	} else if len(os.Args) > 4 { // Ensure there are at least 5 arguments (os.Args[0] to os.Args[4])
		if os.Args[len(os.Args)-1] != "" && (strings.ToLower(os.Args[len(os.Args)-1]) == "shadow" || strings.ToLower(os.Args[len(os.Args)-1]) == "standard" || strings.ToLower(os.Args[len(os.Args)-1]) == "thinkertoy") {
			style = strings.ToLower(os.Args[len(os.Args)-1])
		} else {
			style = "standard"
		}
	} else {
		fmt.Println("No style detected defaulting to standard")
		style = "standard"
	}


	var colorFlag, colorName, substring, inputString string

	for i, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			colorFlag = arg
			colorName = strings.TrimPrefix(colorFlag, "--color=")
			if i+1 < len(args) && len(args) == 4 && (args[3] == "shadow" || args[3] == "standard" || args[3] == "thinkertoy") {
				substring = args[i+1]
				inputString = args[i+2]
			} else if i+2 < len(args) && len(args) == 4 && args[3] != "shadow" && args[3] != "standard" && args[3] != "thinkertoy" {
				substring = args[i+1]
				inputString = strings.Join(args[i+2:], " ")
			} else if i+1 < len(args) && len(args) == 3 && args[2] != "shadow" && args[2] != "standard" && args[2] != "thinkertoy" {
				substring = args[i+1]
				inputString = strings.Join(args[i+2:], " ")
			} else if i+1 < len(args) && len(args) == 3 && (args[2] == "shadow" || args[2] == "standard" || args[2] == "thinkertoy") {
				substring = ""
				inputString = args[i+1]
			} else if len(args) == 2 {
				fmt.Println("Error: Substring or input string is missing.")
				fmt.Println("EX: go run . --color=red kit \"a king kitten have kit\"")
				inputString = args[i+1]
			} else if  i+2 < len(args) && len(args) > 4 && strings.Contains(args[0],"color") && args[len(args)-1] == "shadow" || args[len(args)-1] == "standard" || args[len(args)-1] == "thinkertoy" {
				substring = args[i+1]
				inputString = strings.Join(args[i+2:len(args)-1]," ")
				fmt.Println(substring)
				fmt.Println(inputString)
			} else {
				if strings.Contains(args[0],"color") {
					substring = args[i+1]
					inputString = strings.Join(args[i+2:]," ")
				}
				
			}
		}
	

	if colorFlag == "" && (args[len(args)-1] == "shadow" || args[len(args)-1] == "standard" || args[len(args)-1] == "thinkertoy") {
		inputString = strings.Join(args[:len(args)-1], " ")
	}  
	if colorFlag == "" && !(args[len(args)-1] == "shadow" || args[len(args)-1] == "standard" || args[len(args)-1] == "thinkertoy") {
		inputString = strings.Join(args, " ")
	}  
	}
	if !strings.Contains(inputString, substring) {
		colorName = "black"
	}
	fileLines := functions.Read(style)
	if len(fileLines) == 0 {
        fmt.Println("Error: Failed to load banner style or file is empty.")
        return
    }
	asciiRep := functions.AsciiRep(fileLines)
	asciiArt := functions.PrintStr(inputString, asciiRep)
	coloredArt := functions.ColorArt(asciiArt, substring, colorName, asciiRep, inputString)

	for _, line := range coloredArt {
		fmt.Println(strings.Join(line, ""))
	}
}
