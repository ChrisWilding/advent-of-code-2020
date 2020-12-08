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

func IsValidPasswordOne(input string) bool {
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

func IsValidPasswordTwo(input string) bool {
	matches := re.FindAllStringSubmatch(input, -1)
	first, _ := strconv.Atoi(matches[0][1])
	first--
	second, _ := strconv.Atoi(matches[0][2])
	second--
	letter, _ := utf8.DecodeRuneInString(matches[0][3])
	password := []rune(matches[0][4])

	return password[first] != password[second] && (password[first] == letter || password[second] == letter)
}

func CountValidPasswords(inputs []string, f func(string) bool) int {
	var valid int
	for _, input := range inputs {
		if f(input) {
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
	fmt.Println(CountValidPasswords(lines, IsValidPasswordOne))
	fmt.Println(CountValidPasswords(lines, IsValidPasswordTwo))
}
