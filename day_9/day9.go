package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const AMM = 9

/* ______               __
  / __/ /_______ ______/ /____
 _\ \/ __/ __/ // / __/ __(_-<
/___/\__/_/  \_,_/\__/\__/___/
*/

var Dir = map[rune]Coord{
	'U': {0, -1}, 'D': {0, 1}, 'L': {-1, 0}, 'R': {1, 0},
}

type Coord struct {
	x int
	y int
}

type Knot struct {
	Index      int
	Tail       bool
	Prev, Next *Knot //Prev == towards Tail, Next == Towards Head
	Coord      Coord
}

func main() {
	start := time.Now()
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	dict := map[Coord]struct{}{}
	rope := tie()
	for s.Scan() {
		txt := string(s.Text())
		var dir rune
		var amm int
		fmt.Sscanf(txt, "%c %d", &dir, &amm)
		for i := 0; i < amm; i++ {
			rope.Coord = rope.Add(Dir[dir])
			dict = rope.Prev.Move(dict, amm, dir)
		}

	}
	fmt.Println(time.Since(start))
	fmt.Println("Task 2: ", len(dict))
}

/* __ __          __    ___              __  _
  / //_/__  ___  / /_  / _/_ _____  ____/ /_(_)__  ___  ___
 / ,< / _ \/ _ \/ __/ / _/ // / _ \/ __/ __/ / _ \/ _ \(_-<
/_/|_/_//_/\___/\__/ /_/ \_,_/_//_/\__/\__/_/\___/_//_/___/
*/

func tie() *Knot {
	head := &Knot{
		Index: AMM,
		Tail:  true,
		Coord: Coord{
			x: 0,
			y: 0,
		},
		Prev: nil}

	for i := 1; i <= AMM; i++ {
		n := &Knot{
			Index: AMM - i,
			Coord: Coord{
				x: 0,
				y: 0,
			},
			Prev: head}
		head.Next = n
		head = n
	}
	return head
}

func (k *Knot) Move(dict map[Coord]struct{}, amm int, dir rune) map[Coord]struct{} {
	if dist := k.Next.Sub(k.Coord); Abs(dist.x) > 1 || Abs(dist.y) > 1 {
		k.Coord = k.Add(Coord{Zwrt(dist.x), Zwrt(dist.y)})
	}
	if k.Prev == nil {
		dict[k.Coord] = struct{}{}
	} else {
		dict = k.Prev.Move(dict, amm, dir)
	}
	return dict
}

func (k Knot) Touching() bool {
	arg := k.Sub(k.Coord)
	if Abs(arg.x) > 1 || Abs(arg.y) > 1 {
		return false
	}
	return true
}

func (k Knot) Add(arg Coord) Coord {
	k.Coord.x += arg.x
	k.Coord.y += arg.y
	return k.Coord
}

func (k Knot) Sub(arg Coord) Coord {
	k.Coord.x -= arg.x
	k.Coord.y -= arg.y
	return k.Coord
}

func Abs(arg int) int {
	if arg < 0 {
		return -arg
	}
	return arg
}
func Zwrt(arg int) int {
	if arg < 0 {
		return -1
	} else if arg > 0 {
		return 1
	}
	return 0
}
