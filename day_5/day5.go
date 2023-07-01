package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const ALPHA = "abcdefghijklmnopqrstuv"

func main() {
	//Get crate layout
	f, err := os.Open("crane.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	crates := [][]byte{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := s.Text()
		crates = append(crates, []byte(txt))
	}

	f, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	type Mane struct {
		Qty  int
		From int
		To   int
	}

	s = bufio.NewScanner(f)
	for s.Scan() {
		//Get order
		m := Mane{}
		txt := s.Text()
		arr := strings.Split(txt, " ")

		for i, j := range arr {
			if _, err := strconv.Atoi(j); err == nil {
				switch i {
				case 1:
					m.Qty, err = strconv.Atoi(j)
				case 3:
					m.From, err = strconv.Atoi(j)
					m.From--
				case 5:
					m.To, err = strconv.Atoi(j)
					m.To--
				}
			}
		}
		//Move the crates
		l := len(crates[m.From])
		x := l - m.Qty
		load := []byte{}

		for i := l - 1; i >= x && i >= 0; i-- {
			load = append(load, crates[m.From][i])
			crates[m.From] = crates[m.From][:i]
		}
		crates[m.To] = append(crates[m.To], load...)
		fmt.Println(txt)
		for n, i := range crates {
			fmt.Println(n+1, " ", string(i), " ", len(i))
		}
	}

}
