package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	n := 1
	s := bufio.NewScanner(f)

	//For Task #1 the sollution is the same, except the array length
	buff := []string{}
	for i := 0; i < 14; i++ {
		buff = append(buff, s.Text())
	}
	dupe := false
	s.Split(bufio.ScanRunes)
	for s.Scan() && !dupe {
		txt := s.Text()
		buff = append(buff[1:], txt)
		for i, j := range buff {
			for k, l := range buff {
				if i != k && j == l || l == "" {
					dupe = true
					break
				}
			}
			if dupe {
				break
			}
		}
		if !dupe {
			break
		}
		dupe = false
		n++
	}
	fmt.Println("Task #1: ", n)
}
