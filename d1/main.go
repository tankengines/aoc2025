package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	position int
	zeroes   int
}

func solve(s *Solution, l string) {
	split := strings.Split(l, "")
	right := split[0] == "R"
	ticks, err := strconv.Atoi(strings.Join(split[1:], ""))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("////")
	fmt.Printf("Instruction: %s\n", l)
	fmt.Printf("Ticks: %d\n", ticks)
	fmt.Printf("Start Position: %d\n", s.position)

	if right {
		s.position += ticks
	} else {
		s.position -= ticks
	}

	fmt.Printf("Preround position: %d\n", s.position)

	revolutions := 0
	revolutions = int(math.Floor(float64(s.position) / 100))
	if revolutions < 0 {
		revolutions = -revolutions
	}

	s.zeroes += revolutions
	fmt.Printf("Zeroes: %d\n", s.zeroes)

	if s.position > 99 {
		s.position = s.position % 100
	}

	if s.position < 0 {
		s.position = -(100 + s.position) % 100

		if s.position < 0 {
			s.position = -s.position
		}
	}

	fmt.Printf("End Position: %d\n", s.position)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	sln := Solution{
		position: 50,
		zeroes:   0,
	}

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

	fmt.Printf("%d\n", sln.zeroes)
}
