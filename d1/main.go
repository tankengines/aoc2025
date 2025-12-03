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
	fmt.Printf("Zeroes: %d\n", s.zeroes)
	fmt.Printf("Instruction: %s\n", l)
	fmt.Printf("Start Position: %d\n", s.position)

	i := 1

	for {
		if right {
			s.position += 1
		} else {
			s.position -= 1
		}

		if s.position == -1 {
			s.position = 99
		}

		if s.position == 100 {
			s.position = 0
		}

		if s.position == 0 {
			s.zeroes += 1
		}

		if i == ticks {
			break
		}

		i++
	}

	fmt.Printf("End Position: %d\n", s.position)

	return

	// Absolutely atrocious attempt to do this algorithmically... keep it simple, stupid!
	start := s.position

	if right {
		s.position += ticks
	} else {
		s.position -= ticks
	}

	fmt.Printf("Preround position: %d\n", s.position)

	// edge case 1: old pointer at 0, move left, gets counted
	// 2: new pointer at 0, but moved 99 ticks
	revolutions := 0.0
	revolutions = math.Floor(math.Abs(float64(s.position)) / 100)

	if s.position > 99 {
		s.position = s.position % 100
	}

	if s.position < 0 {
		s.position = -(100 + s.position) % 100

		if s.position < 0 {
			s.position = -s.position
		}
	}

	if start != 0 && s.position < 0 {
		revolutions += 1
	}

	if s.position == 0 {
		revolutions += 1
	}

	s.zeroes += int(revolutions)

	fmt.Printf("Rounded rev: %d\n", int(revolutions))

	fmt.Printf("Zeroes: %d\n", s.zeroes)

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
