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

func main() {
	start := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var n, m int
	for s.Scan() {
		txt := s.Text()
		matrix := [][]int{
			{0, 0}, {0, 0},
		}
		elfs := strings.Split(txt, ",")
		for i := range elfs {
			arr := strings.Split(elfs[i], "-")
			matrix[i][0], err = strconv.Atoi(arr[0])
			matrix[i][1], err = strconv.Atoi(arr[1])
		}
		switch {
		case matrix[0][0] <= matrix[1][0] && matrix[0][1] >= matrix[1][1]:
			n++
		case matrix[0][0] >= matrix[1][0] && matrix[0][1] <= matrix[1][1]:
			n++
		case matrix[0][1] >= matrix[1][0] && matrix[0][0] <= matrix[1][1]:
			m++
		case matrix[0][1] <= matrix[1][0] && matrix[0][0] >= matrix[1][1]:
			m++
		default:
		}
	}
	fmt.Println(n)
	fmt.Println(n + m)
	fmt.Println(time.Since(start))
}
