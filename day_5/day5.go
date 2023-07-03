package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Crane struct {
	Qty  int
	From int
	To   int
	Mat  [][]byte
}

const ALPHA = "abcdefghijklmnopqrstuv"

func main() {
	//Get crate layout
	f, err := os.Open("crane.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	otpt1, otpt2 := Crane{}, Crane{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := s.Text()
		otpt1.Mat = append(otpt1.Mat, []byte(txt))
		otpt2.Mat = append(otpt2.Mat, []byte(txt))
	}

	f, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s = bufio.NewScanner(f)
	for s.Scan() {
		//Get order
		txt := s.Text()
		arr := strings.Split(txt, " ")

		for i, j := range arr {
			if _, err := strconv.Atoi(j); err == nil {
				switch i {
				case 1:
					otpt1.Qty, _ = strconv.Atoi(j)
					otpt2.Qty, _ = strconv.Atoi(j)
				case 3:
					otpt1.From, _ = strconv.Atoi(j)
					otpt2.From, _ = strconv.Atoi(j)
					otpt1.From--
					otpt2.From--
				case 5:
					otpt1.To, _ = strconv.Atoi(j)
					otpt2.To, _ = strconv.Atoi(j)
					otpt1.To--
					otpt2.To--
				}
			}
		}
		otpt1.pop()
		otpt2.appnd()
	}

	fmt.Println("\n --- Task #1 ---")
	fmt.Println()
	for n, i := range otpt1.Mat {
		fmt.Println(n+1, " ", string(i), " ", len(i))
	}

	fmt.Println("\n --- Task #2 ---")
	fmt.Println()
	for n, i := range otpt2.Mat {
		fmt.Println(n+1, " ", string(i), " ", len(i))
	}
}

func (c *Crane) pop() {
	//Move the crates using push/pop mechanism
	l := len(c.Mat[c.From])
	x := l - c.Qty
	load := []byte{}

	for i := l - 1; i >= x && i >= 0; i-- {
		load = append(load, c.Mat[c.From][i])
		c.Mat[c.From] = c.Mat[c.From][:i]
	}
	c.Mat[c.To] = append(c.Mat[c.To], load...)
}
func (c *Crane) appnd() {
	//Move the crates by reattaching subarrays
	l := len(c.Mat[c.From])
	x := l - c.Qty
	load := c.Mat[c.From][x:]
	c.Mat[c.To] = append(c.Mat[c.To], load...)
	c.Mat[c.From] = append(c.Mat[c.From][:x])

}
