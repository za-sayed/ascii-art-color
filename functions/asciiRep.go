package functions

func AsciiRep(fileLines []string) [][]string {
	var asciiRep [][]string
	var arr []string
	counter := 0
	// counter is incremented by the number of lines in filelines
	for _, line := range fileLines {
		// If counter starts again skip the first line in the standard text file which seperates each ASCII value
		if counter == 0 {
			counter++
			continue
		}
		// every line is appended seperately into a string slice called arr
		arr = append(arr, line)
		// Every time a new line is appended to arr the counter is incremented by 1
		counter++
		// Once the counter reaches 9 (because each character is 8 lines followed by 1 space)
		if counter == 9 {
			// append that character array to asciiRep
			asciiRep = append(asciiRep, arr)
			// Reset arr and counter values to 0 for when next word begins
			arr = nil
			counter = 0
		}
	}
	return asciiRep
}
