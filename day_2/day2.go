package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	start := time.Now()
	dict1 := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}
	scores := [][]int{
		{4, 1, 7},
		{8, 5, 2},
		{3, 9, 6},
	}
	//Determine what to choose in a current round, depending on expected outcome
	task2 := [][]int{
		{2, 0, 1},
		{0, 1, 2},
		{1, 2, 0},
	}

	// Read the text file containing the list
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//a, b := []string{}, []string{}
	s := bufio.NewScanner(f)
	score1, score2 := 0, 0
	for s.Scan() {
		txt := s.Text()
		x := dict1[string(txt[len(txt)-1])]
		y := dict1[string(txt[0])]
		z := task2[y][x]
		score1 += scores[x][y]
		score2 += scores[z][y]
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Task #1: ", score1)
	fmt.Println("Task #2: ", score2)
	fmt.Println(time.Since(start))
}
