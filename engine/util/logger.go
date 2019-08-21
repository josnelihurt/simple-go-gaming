package util

import "fmt"

var Logger chan string
func DoLog(input <-chan string) {
	for line := range input {
		fmt.Println(line)
	}
}
