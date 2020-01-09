package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("input a string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)
	match := re.FindString(text)
	fmt.Println(match)
}
