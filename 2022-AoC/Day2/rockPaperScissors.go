package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 0 lost, 3 draw, 6 won
// A Rock, B Paper, C Scissors
// X Rock, Y Paper, Z Scissors
// x - 0 tie, -1 win, -2 lose
// y - 0 tie, 1 win, -1 lose
// z - 0 tie, 1 win, 2 lose
var results = map[string]int{
	"AX": 4,
	"BX": 1,
	"CX": 7,
	"AY": 8,
	"BY": 5,
	"CY": 2,
	"AZ": 3,
	"BZ": 9,
	"CZ": 6,
}

// A Rock, B Paper, C Scissors
// X Rock, Y Paper, Z Scissors
// X lose, Y draw, Z win
var combo = map[string]string{
	"AX": "AZ",
	"BX": "BX",
	"CX": "CY",
	"AY": "AX",
	"BY": "BY",
	"CY": "CZ",
	"AZ": "AY",
	"BZ": "BZ",
	"CZ": "CX",
}

func clean(line string) string{
	return strings.ReplaceAll(strings.TrimSuffix(line, "\n"), " ", "")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sum := 0
	
	for {
	    line, err := reader.ReadString('\n')
	    if err != nil {
	        break
	    }
        
		sum += results[combo[clean(line)]]
	}
	fmt.Print(sum)
}
