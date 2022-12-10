package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//     [M]             [Z]     [V]
//     [Z]     [P]     [L]     [Z] [J]
// [S] [D]     [W]     [W]     [H] [Q]
// [P] [V] [N] [D]     [P]     [C] [V]
// [H] [B] [J] [V] [B] [M]     [N] [P]
// [V] [F] [L] [Z] [C] [S] [P] [S] [G]
// [F] [J] [M] [G] [R] [R] [H] [R] [L]
// [G] [G] [G] [N] [V] [V] [T] [Q] [F]
//  1   2   3   4   5   6   7   8   9

type supply []string

type supplyStacks struct {
	supply []supply
}

func (s *supplyStacks) Pop(si int, n int) []string {
	curSupLen := len(s.supply[si])
	items := s.supply[si][curSupLen-n:]
	s.supply[si] = s.supply[si][:curSupLen-n]
	return items
}

func (s *supplyStacks) Push(si int, sL []string) {
	s.supply[si] = append(s.supply[si], sL...)
}

func (s *supplyStacks) Insert(si int, c string) {
	s.supply[si] = append(supply{c}, s.supply[si]...)
	return
}

func initStacks(n int) *supplyStacks {
	s := supplyStacks{
		supply: make([]supply, n),
	}

	for i := 0; i < n; i++ {
		s.supply[i] = make(supply, 0)
	}
	return &s
}

func (s *supplyStacks) loadStack(r *bufio.Reader, count int) int {
	var t = 0
	line, err := r.ReadString('\n')
	if err != nil {
	    os.Exit(1)
	}
	line = strings.TrimSuffix(line, "\n")
	if len(line) == 0 {
	    return 2
	}

	for i := 0; i < count; i++ {
		c := string(line[i*4+1])
		if t == 0 {

			if _, err := strconv.Atoi(string(c)); err == nil {
				return 1
			}
		}
		t++
		if c != " " {
    		s.Insert(i, c)  
		} 

	}
    
	return 0
}

func (s supplyStacks) getTopString() string {
	var b []string
	for _, s := range s.supply {
		b = append(b, s[len(s)-1])
	}

	return strings.Join(b, "")
}

func (s supplyStacks) move(count, src, dest int) {
    fmt.Println(s.supply[src])
    fmt.Println(s.supply[dest])
    items := s.Pop(src, count)
    fmt.Println("moving ", items, " to ", dest, s.supply[dest])
	s.Push(dest, items)
    fmt.Println(s.supply[src])
    fmt.Println(s.supply[dest])
}

func main() {
	ss := initStacks(9)

	reader := bufio.NewReader(os.Stdin)
	exitCode := 0
	for exitCode != 2 {
		exitCode = ss.loadStack(reader, 9)
	}
	
	fmt.Println(ss)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(ss.getTopString())
			os.Exit(0)
		}
		

		line = strings.TrimSpace(line)
		var count, src, dest int
		fmt.Sscanf(line, "move %d from %d to %d", &count, &src, &dest)
		fmt.Println(count, src, dest)
		fmt.Println(ss.supply)
		ss.move(count, src-1, dest-1)
	}

}
