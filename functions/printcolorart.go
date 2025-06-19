package functions

import (
	"fmt"
	"strings"
)

// PrintStr function converts the input string to its corresponding ASCII art representation as a 2d stringg slice
func PrintStr(inputString string, asciiRep [][]string) [][]string {
	charHeight := len(asciiRep[0])
	output := make([][]string, charHeight)

	for i := range output {
		output[i] = make([]string, 0, len(inputString))
	}

	for _, char := range inputString {
		if char >= 32 && char <= 126 {
			index := int(char) - 32
			if index < 0 || index >= len(asciiRep) {
				fmt.Printf("Warning: ASCII representation for character '%c' not found.\n", char)
				continue
			}
			for i := 0; i < charHeight; i++ {
				output[i] = append(output[i], asciiRep[index][i])
			}
		} else {
			for i := 0; i < charHeight; i++ {
				output[i] = append(output[i], " ") // Handle non-printable characters
			}
		}
	}
	return output
}

func ColorArt(asciiArt [][]string, substring string, colorName string, asciiRep [][]string, inputString string) [][]string {
	colorCode := getColorCode(colorName)
	resetCode := "\033[0m" //black
	if colorCode == "" {
		fmt.Printf("Error: Invalid color name: %s\n", colorName)
		fmt.Println("Usage: go run . --color=red [STRING]")
		return nil
	}
	// Generate ASCII art for the substring
	coloredArt := make([][]string, len(asciiArt))
	//substringArt := PrintStr(substring, asciiRep)
	for i := range asciiArt {
		coloredArt[i] = make([]string, len(asciiArt[i]))
		copy(coloredArt[i], asciiArt[i])
		if len(substring) == 1 { // Handle single-character substring
			for x := 0; x < len(inputString); x++ {
				// find the index of where the substring is postitioned in the inputstring.
				if string(inputString[x]) == substring {
					// Then change the color of coloredArt[i] only  at that index
					coloredArt[i][x] = colorCode + coloredArt[i][x] + resetCode
				}
			}
		}
		// handle two character strings
		if len(substring) == 2 {
			for x := 0; x < len(inputString); x++ {
				// find the index of where the substring is postitioned in the inputstring.
				if x+1 < len(inputString) && string(inputString[x]) + string(inputString[x+1]) == string(substring[0]) + string(substring[1]) {
						coloredArt[i][x] = colorCode + coloredArt[i][x] + resetCode
						coloredArt[i][x+1] = colorCode + coloredArt[i][x+1] + resetCode	
						x++ // Move the index forward to avoid overlapping matches
				}
			}
		}
		// Handle multi-character strings (length greater than 2)
if len(substring) > 2 {
    for x := 0; x <= len(inputString)-len(substring); x++ {
        // find the index of where the substring is positioned in the inputString.
        if string(inputString[x:x+len(substring)]) == substring {
            // Apply color to the corresponding range in coloredArt
            for y := 0; y < len(substring); y++ {
                coloredArt[i][x+y] = colorCode + coloredArt[i][x+y] + resetCode
            }
			   // Move the index forward to avoid overlapping matches. 
			   x += len(substring) - 1
        }
    }
}
	}
		return coloredArt
}

// getColorCode returns code for the specified color name
func getColorCode(colorName string) string {
	switch strings.ToLower(colorName) {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "magenta":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	case "black":
		return "\033[30m"
	case "orange":
		return "\033[38;5;214m"
	default:
		return ""
	}
}
