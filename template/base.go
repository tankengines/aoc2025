package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Solution struct {
}

func solve(s *Solution, l string) {
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	sln := Solution{}

	for {
		l, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Fatal error reading stdin: %s\n", err)
		}

		l = strings.TrimSpace(l)

		if len(l) == 0 {
			continue
		}

		solve(&sln, l)
	}

	// fmt.Printf("Solution\n", sln.)
}
