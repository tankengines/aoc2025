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
	invalids []int
}

func solve(s *Solution, l string) {
	// Remove trailing comma
	l = l[:len(l)-1]

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

		str := strconv.Itoa(i)

		repeat := isAllRepeat(str)
		if repeat {
			s.invalids = append(s.invalids, i)
			fmt.Printf("%d\n", i)
			i++
			continue
		}

		if str[:len(str)/2] == str[len(str)/2:] {
			s.invalids = append(s.invalids, i)
			fmt.Printf("%d\n", i)
		}

		i++
	}
}

func isAllRepeat(s string) bool {
	isSame := true
	if len(s) == 1 {
		return true
	}

	for i := range len(s) - 1 {
		if s[i] == s[i+1] {
			continue
		}
		isSame = false
		break
	}

	if len(s)%2 == 1 {
		return false
	}

	return isSame
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
	n := 0
	for i := 0; i < len(sln.invalids); i++ {
		n += sln.invalids[i]
	}

	fmt.Printf("Final sum: %d\n", n)
}
