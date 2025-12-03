package main

import (
	"bufio"
	"fmt"
	// "io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	invalids []int64
}

func solve(s *Solution, l string) {
	fromTo := strings.Split(l, "-")

	from, err := strconv.ParseInt(fromTo[0], 10, 64)

	if err != nil {
		log.Fatalf("Fatal error doing atoi: %s\n", err)
	}

	to, err := strconv.ParseInt(fromTo[1], 10, 64)

	if err != nil {
		log.Fatalf("Fatal error doing atoi: %s\n", err)
	}

	for i := from; i <= to; i++ {
		str := strconv.FormatInt(i, 10)

		repeat := isAllRepeat(str)
		if repeat {
			s.invalids = append(s.invalids, i)
			// fmt.Printf("%d\n", i)
			continue
		}

		if str[:len(str)/2] == str[len(str)/2:] {
			s.invalids = append(s.invalids, i)
			// fmt.Printf("%d\n", i)
			continue
		}
	}
}

func isAllRepeat(s string) bool {
	if len(s) == 1 {
		return false
	}

	if (len(s) % 2) == 1 {
		return false
	}

	isSame := true

	for i := range len(s) - 1 {
		if s[i] == s[i+1] {
			continue
		}
		isSame = false
		break
	}

	return isSame
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	sln := Solution{invalids: []int64{}}

	text, _ := reader.ReadString('\n')

	list := strings.Split(text, ",")
	for i := 0; i < len(list); i++ {
		// Remove trailing whitespace
		l := strings.Trim(list[i], "\n")

		fmt.Println(l)
		solve(&sln, l)
	}

	var n int64 = 0
	for i := 0; i < len(sln.invalids); i++ {
		n += sln.invalids[i]
	}

	fmt.Printf("Final sum: %d\n", n)
}
