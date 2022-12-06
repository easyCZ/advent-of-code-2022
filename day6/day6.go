package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	signal := parse(os.Stdin)
	fmt.Println(signal)

	fmt.Println("Packet Marker", findMarker(signal))
	fmt.Println("Message Marker", findMessageMarker(signal))
}

func findMarker(s string) int {
	seen := map[string]struct{}{}

	for i := 0; i < 4; i++ {
		seen[string(s[i])] = struct{}{}
	}

	for i := 4 + 1; i < len(s); i++ {
		item := string(s[i])

		if _, ok := seen[item]; !ok {
			cs := counts(s[i-4 : i])

			if len(cs) == 4 {
				return i
			}
		}

		seen[string(s[i])] = struct{}{}
	}

	return -1
}

func findMessageMarker(s string) int {
	size := 14
	for i := size; i < len(s); i++ {
		cs := counts(s[i-size : i])

		if len(cs) == size {
			return i
		}
	}

	return -1
}

func parse(r io.Reader) string {
	s := bufio.NewScanner(r)

	s.Scan()
	t := strings.TrimSpace(s.Text())
	return t
}

func counts(s string) map[string]int {
	cs := map[string]int{}

	for _, c := range s {
		if _, ok := cs[string(c)]; ok {
			cs[string(c)] += 1
		} else {
			cs[string(c)] = 1
		}
	}

	return cs
}
