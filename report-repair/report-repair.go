package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func read(path string) []int {
	var numbers []int
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return numbers
}

func one(numbers []int) int {
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1+n2 == 2020 {
				return n1 * n2
			}
		}
	}
	return 0
}

func two(numbers []int) int {
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			for _, n3 := range numbers {
				if n1+n2+n3 == 2020 {
					return n1 * n2 * n3
				}
			}
		}
	}
	return 0
}

func main() {
	path := os.Args[1]
	numbers := read(path)
	fmt.Println(one(numbers))
	fmt.Println(two(numbers))
}
