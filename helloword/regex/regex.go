package main

import (
	"fmt"
	"regexp"
)

const text = "My email is wqq@gmail.com"

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)
	match := re.FindAllString(text, -1)
	fmt.Println(match)
}
