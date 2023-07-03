package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const LIMIT = 100000

type Dir struct {
	Size     int
	Files    []int
	Embedded bool
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	repo := [][]Dir{}
	depth, width := 0, 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := strings.Split(s.Text(), " ")
		switch txt[1] {
		case "cd":
			switch txt[2] {
			case "/":
				depth = 0
				width++
			case "..":
				depth--
				width = len(Dir[depth])
			default:
				depth++
			}
		case "ls":
		default:
			switch txt[0] {

			}

		}
	}
	sum := 0
	for _, i := range arr {
		if i.Size <= LIMIT {
			sum += i.Size
		}
	}
	fmt.Println(sum)
}
