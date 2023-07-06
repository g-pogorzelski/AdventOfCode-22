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

const (
	LIMIT  = 100000
	UPDATE = 30000000
	DISK   = 70000000
)

type Dir struct {
	Name    string
	Size    int
	PrevDir *Dir
	SubDir  []*Dir
}

// add weights of subdirectories
func (d *Dir) weight() {
	for _, j := range d.SubDir {
		j.weight()
		j.PrevDir.Size += j.Size
	}
}

func main() {
	start := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	repo, head := Dir{}, &Dir{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := strings.Split(s.Text(), " ")
		switch txt[1] {
		case "cd":
			switch txt[2] {
			case "/":
				repo = Dir{PrevDir: nil, Name: txt[2]}
				head = &repo
			case "..":
				head.PrevDir.SubDir = append(head.PrevDir.SubDir, head)
				head = head.PrevDir
			default:
				head = &Dir{
					Name:    txt[2],
					PrevDir: head,
				}
			}
		case "ls":
		default:
			if n, err := strconv.Atoi(txt[0]); err == nil {
				head.Size += n
			}
		}
	}
	for head.PrevDir != nil {
		head.PrevDir.SubDir = append(head.PrevDir.SubDir, head)
		head = head.PrevDir
	}
	lim := getAll(repo)
	fmt.Println("Task #1: ", lim)
	repo.weight()
	req := UPDATE - (DISK - repo.Size)
	arr := findFit(repo, req)
	args := Sort(arr)
	fmt.Println("Task #2: ", args[len(args)-1])
	fmt.Println(time.Since(start))
}

func getAll(d Dir) (size int) {
	for _, j := range d.SubDir {
		arg := getAll(*j)
		size += arg
		d.Size += j.Size
	}
	if d.Size <= LIMIT {
		size += d.Size
	}
	return
}

func findFit(d Dir, req int) (otpt []int) {
	for _, j := range d.SubDir {
		if j.Size-req > 0 {
			otpt = append(otpt, findFit(*j, req)...)
		}

	}
	if d.Size-req > 0 {
		otpt = append(otpt, d.Size)
	}
	return
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
		n := (l / 2) + 1
		a := Sort(arr[:n])
		b := Sort(arr[n:])
		i, j := 0, 0
		for i < len(a) && j < len(b) {
			if a[i] > b[j] {
				otpt = append(otpt, a[i])
				i++
			} else {
				otpt = append(otpt, b[j])
				j++
			}
		}
		if i > j && len(b) > j {
			otpt = append(otpt, Sort(b[j:])...)
		} else {
			otpt = append(otpt, Sort(a[i:])...)
		}
		return otpt
	}
}
