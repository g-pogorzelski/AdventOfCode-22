package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	// Read the text file containing the list
	f, err := os.Open("list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var elfs []int
	var elf int
	s := bufio.NewScanner(f)

	for s.Scan() {
		// if the scanned line is transferable into int, add it to current elf
		if arg, err := strconv.Atoi(s.Text()); err != nil {
			elfs = append(elfs, elf)
			elf = 0
		} else {
			elf += arg
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	x := Sort(elfs)
	fmt.Println(x[1] + x[2] + x[3])
	fmt.Println(time.Since(start))
}

func Sort(arr []int) []int {
	otpt := []int{}
	switch l := len(arr); l {
	case 1:
		return arr
	case 2:
		if arr[0] < arr[1] {
			return []int{arr[1], arr[0]}
		}
		return []int{arr[0], arr[1]}
	default:
		n := l / 2
		a := Sort(arr[:n])
		b := Sort(arr[n:])
		i, j := 0, 0
		for i < len(a) && j < len(b) {
			if a[i] < b[j] {
				otpt = append(otpt, b[j])
				j++
			} else if i <= j {
				otpt = append(otpt, a[i])
				i++
			} else {
				otpt = append(otpt, b[j])
				j++
			}
		}
		if i > j {
			otpt = append(otpt, b[j:]...)
		} else {
			otpt = append(otpt, a[i:]...)
		}
		return otpt
	}
}
