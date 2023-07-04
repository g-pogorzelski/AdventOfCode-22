package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const LIMIT = 100000

type Dir struct {
	Name    string
	Size    int
	PrevDir *Dir
	SubDir  []Dir
}

func main() {
	f, err := os.Open("testInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	repo, head := Dir{}, Dir{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		txt := strings.Split(s.Text(), " ")
		switch txt[1] {
		case "cd":
			switch txt[2] {
			case "/":
				repo = Dir{PrevDir: nil, Name: txt[2]}
				head = repo
			case "..":
				if head.Name != "/" {
					head = *head.PrevDir
				}
			default:
				if head.PrevDir != nil {
					head = Dir{PrevDir: &repo.SubDir[len(repo.SubDir)-1], Name: txt[2]}
					head.PrevDir.SubDir = append(head.PrevDir.SubDir, head)
				} else {
					head = Dir{PrevDir: &repo, Name: txt[2]}
					repo.SubDir = append(repo.SubDir, head)
				}
			}
		case "ls":
		default:
			switch n, err := strconv.Atoi(txt[0]); err {
			case nil:
				head.Size += n
			default:

			}
		}
	}
	sum := 0
	fmt.Println(sum)
}
