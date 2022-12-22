package main

import (
	"fmt"
	"regexp"
)

func insertCommaSpaces(s string) string {
	re := regexp.MustCompile(",[^ ]")
	return re.ReplaceAllString(s, ", ")
}
func main() {
	input := "Q1 months are January,February,March."
	output := insertCommaSpaces(input)
	fmt.Println(output)

}
