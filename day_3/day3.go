package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const ALPHABETA = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	matrix := [][]string{}
	// Read the text file containing the list
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	//Task #1
	var n, score int
	for s.Scan() {
		//Gather inputs in groups of 3 for Task #2
		if n%3 == 0 {
			matrix = append(matrix, []string{})
		}
		txt := s.Text()
		matrix[n/3] = append(matrix[n/3], txt)
		l := len(s.Text()) / 2
		a := string(txt[:l])
		b := string(txt[l:])
		arg := findCommon(a, b)
		for i := range ALPHABETA {
			if byte(ALPHABETA[i]) == arg {
				score += i + 1
				break
			}
		}
		n++
	}
	fmt.Println("Task 1: ", score)
	//Task #2
	score = 0
	for i := range matrix {
		score += findBadge(matrix[i])
	}
	fmt.Println(score)
}

func findCommon(a, b string) byte {
	var arg byte
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				arg = a[i]
				break
			}
		}
	}
	return arg
}

func findBadge(arr []string) int {
	x, y, z := arr[0], arr[1], arr[2]
	var arg byte
	common := []byte{}
	for i := 0; i < len(x); i++ {
		a := string(x[i])
		for j := 0; j < len(y); j++ {
			b := string(y[j])
			if a == b && uniq(common, x[i]) {
				common = append(common, x[i])
			}

		}
	}
	for i := 0; i < len(common); i++ {
		for j := 0; j < len(z); j++ {
			if common[i] == z[j] {
				arg = z[j]
				break
			}
		}
	}
	for i := range ALPHABETA {
		if byte(ALPHABETA[i]) == arg {
			return i + 1
		}
	}
	return 0
}

func uniq(arr []byte, arg byte) bool {
	for i := range arr {
		if arr[i] == arg {
			return false
		}
	}
	return true
}
