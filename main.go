package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Cannot open input file: %s\n", err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Cannot read input file: %s\n", err)
	}

	input := strings.Split(string(b), "\n")

	current := 50
	zeroes := 0
	for i := 0; i < len(input); i++ {
		if len(input[i]) == 0 {
			continue
		}

		split := strings.Split(input[i], "")
		d := split[0] == "R"
		count, err := strconv.Atoi(strings.Join(split[1:], ""))

		if err != nil {
			log.Fatal(err)
		}

		if d {
			current = current + count
		} else {
			current = (current - count)
		}

		if current < 0 || current > 99 {
			current = current % 100
		}

		if current == 0 {
			zeroes++
		}
	}

	fmt.Printf("%d\n", zeroes)
}
