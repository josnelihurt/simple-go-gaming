package util

import "fmt"

// Logger fixme
var Logger chan string

// DoLog fixme
func DoLog(input <-chan string) {
	for line := range input {
		fmt.Println(line)
	}
}
