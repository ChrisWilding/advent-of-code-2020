package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var re = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

func IsValidPassword(input string) bool {
	matches := re.FindAllStringSubmatch(input, -1)
	lower, _ := strconv.Atoi(matches[0][1])
	upper, _ := strconv.Atoi(matches[0][2])
	letter, _ := utf8.DecodeRuneInString(matches[0][3])
	password := matches[0][4]

	var count int
	for _, r := range password {
		if r == letter {
			count++
		}
	}

	return count >= lower && count <= upper
}

func CountValidPasswords(inputs []string) int {
	var valid int
	for _, input := range inputs {
		if IsValidPassword(input) {
			valid++
		}
	}
	return valid
}

func read(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func main() {
	path := os.Args[1]
	lines := read(path)
	fmt.Println(CountValidPasswords(lines))
}
