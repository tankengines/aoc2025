package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
}

func solve(s *Solution, l string) {
	// Remove trailing comma
	l = l[:len(l)-1]

	fmt.Printf("%s\n", l)

	fromTo := strings.Split(l, "-")

	from, err := strconv.Atoi(fromTo[0])

	if err != nil {
		log.Fatalf("Fatal error doing atoi: %s\n", err)
	}

	to, err := strconv.Atoi(fromTo[1])

	if err != nil {
		log.Fatalf("Fatal error doing atoi: %s\n", err)
	}

	i := from

	for {
		if i > to {
			break
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	sln := Solution{}

	for {
		l, err := reader.ReadString(',')

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
