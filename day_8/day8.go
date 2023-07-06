package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cross struct {
	U  []int
	D  []int
	L  []int
	R  []int
	cU bool
	cD bool
	cL bool
	cR bool
}

func main() {
	start := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	mat := [][]int{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := strings.Split(s.Text(), "")
		mat = append(mat, []int{})
		l := len(mat) - 1
		for _, i := range txt {
			arg, _ := strconv.Atoi(i)
			mat[l] = append(mat[l], arg)
		}
	}
	/*
		All the edges are visible, therefore the default ammount of visible trees
		should be the sum of those that sit on the edges of the grid
	*/
	vis := 2*len(mat) + 2*len(mat[0]) - 4

	hiScore := 0
	for i, j := 1, 1; i+1 < len(mat); j++ {
		c := Cross{}
		arg := mat[i][j]
		for n, k := range mat {
			switch {
			case n < i:
				c.U = append(c.U, k[j])
			case n > i:
				c.D = append(c.D, k[j])
			default:
				c.L = append(c.L, k[:j]...)
				c.R = append(c.R, k[j+1:]...)
			}
		}
		//Count the scenery score, starting from the closest points on the grid
		score := []int{
			sceneryDsc(arg, c.U),
			sceneryDsc(arg, c.L),
			sceneryAsc(arg, c.D),
			sceneryAsc(arg, c.R),
		}
		scene := score[0] * score[1] * score[2] * score[3]
		if scene > hiScore {
			hiScore = scene
		}

		c.cU = isCovered(arg, c.U)
		c.cD = isCovered(arg, c.D)
		c.cL = isCovered(arg, c.L)
		c.cR = isCovered(arg, c.R)
		if !c.cU || !c.cD || !c.cL || !c.cR {
			vis++
		}

		// Mechanism for moving within the grid's edges
		if j+2 == len(mat[0]) {
			j = 0
			i++
		}
	}
	fmt.Println("Task #1: ", vis)
	fmt.Println("Task #2: ", hiScore)
	fmt.Println(time.Since(start))
}

func isCovered(arg int, arr []int) bool {
	for _, i := range arr {
		if arg <= i {
			return true
		}
	}
	return false
}

func sceneryAsc(arg int, arr []int) int {
	for i, j := range arr {
		if arg <= j {
			return i + 1
		}
	}
	return len(arr)
}
func sceneryDsc(arg int, arr []int) int {
	l := len(arr)
	for i := range arr {
		x := l - i
		if arg <= arr[x-1] {
			return i + 1
		}
	}
	return l
}
