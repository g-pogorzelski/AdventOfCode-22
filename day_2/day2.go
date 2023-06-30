package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	const A, B, C = 1, 2, 3
	const X, Y, Z = 1, 2, 3

	start := time.Now()
	// Read the text file containing the list
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//a, b := []string{}, []string{}
	s := bufio.NewScanner(f)
	score := 0
	for s.Scan() {
		txt := s.Text()
		a := string(txt[0])
		b := string(txt[len(txt)-1])
		switch {
		case a == "A" && b == "X":
			score += 0 + 3
		case a == "A" && b == "Y":
			score += 3 + 1
		case a == "A" && b == "Z":
			score += 6 + 2
		case a == "B" && b == "X":
			score += 0 + 1
		case a == "B" && b == "Y":
			score += 3 + 2
		case a == "B" && b == "Z":
			score += 6 + 3
		case a == "C" && b == "X":
			score += 0 + 2
		case a == "C" && b == "Y":
			score += 3 + 3
		case a == "C" && b == "Z":
			score += 6 + 1
		default:
			break
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(score)
	fmt.Println(time.Since(start))
}
