package main

import (
	"fmt"
	"os"
	"strings"
)

type CPU struct {
	val  int
	inst int
}

func main() {
	X := CPU{1, 0}

	f, _ := os.ReadFile("input.txt")
	s := strings.Split(string(f), "\n")
	n := X.Read(1, 20, s)

	for i := 20; i < 220; i += 40 {
		n += X.Read(i, (i + 40), s)
	}
	fmt.Println(n)
}

func (X *CPU) Read(start int, finish int, s []string) int {
	var com string
	var arg int
	i := X.inst
	for ; start < finish; i++ {
		fmt.Sscanf(s[i], "%v %d", &com, &arg)
		if com != "noop" {
			start++
			if !(start < finish) {
				break
			}
			X.val += arg
		}
		start++
	}
	X.inst = i
	otpt := X.val * finish
	return otpt
}
