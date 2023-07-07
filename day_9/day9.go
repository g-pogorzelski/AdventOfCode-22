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

const AMM = 10

type Coord struct {
	x int
	y int
}

type Knot struct {
	Tail  bool
	Prev  *Knot
	Coord Coord
}

var Dir = map[string]Coord{
	"R": {
		x: 1,
		y: 0,
	},
	"L": {
		x: -1,
		y: 0,
	},
	"U": {
		x: 0,
		y: 1,
	},
	"D": {
		x: 0,
		y: -1,
	},
}

func main() {
	start := time.Now()
	f, err := os.Open("testInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	rope := Knot{}
	head := &Knot{
		Tail: true,
		Coord: Coord{
			x: 0,
			y: 0,
		},
		Prev: nil}

	for i := 1; i < AMM; i++ {
		n := &Knot{
			Coord: Coord{
				x: 0,
				y: 0,
			},
			Prev: head}
		head = n
	}
	rope = *head
	head = &rope
	dict := make(map[Coord]bool)
	for s.Scan() {
		txt := strings.Split(s.Text(), " ")
		n, err := strconv.Atoi(txt[1])
		if err != nil {
			log.Fatal(err)
		}
		dir := Dir[txt[0]]
		for i := 0; i < n; i++ {
			rope.Coord.x += dir.x
			rope.Coord.y += dir.y
		}
		for head.Prev != nil {
			if head.Prev.Tail {
				dict[head.Prev.Coord] = true
			}
			for !head.Touching() {
				arg := head.Move(txt[0], dir)
				if head.Prev.Tail {
					dict[head.Prev.Coord] = arg
				}
			}
			head = head.Prev
		}

		head = &rope
	}
	fmt.Println(time.Since(start))
	fmt.Println("Task 2: ", len(dict))
}

func (k Knot) Touching() bool {
	arg := k.Prev
	if k.Coord.x > arg.Coord.x+1 || k.Coord.x < arg.Coord.x-1 {
		return false
	} else if k.Coord.y > arg.Coord.y+1 || k.Coord.y < arg.Coord.y-1 {
		return false
	}
	return true
}

func (k *Knot) Move(txt string, arg Coord) bool {
	if txt == "U" && k.Coord.x != k.Prev.Coord.x {
		k.Prev.Coord.x = k.Coord.x
		k.Prev.Coord.y += arg.y
	} else if txt == "D" && k.Coord.x != k.Prev.Coord.x {
		k.Prev.Coord.x = k.Coord.x
		k.Prev.Coord.y += arg.y
	} else if txt == "L" && k.Coord.y != k.Prev.Coord.y {
		k.Prev.Coord.x += arg.x
		k.Prev.Coord.y = k.Coord.y
	} else if txt == "R" && k.Coord.y != k.Prev.Coord.y {
		k.Prev.Coord.x += arg.x
		k.Prev.Coord.y = k.Coord.y
	} else {
		k.Prev.Coord.x += arg.x
		k.Prev.Coord.y += arg.y
	}
	return true
}
