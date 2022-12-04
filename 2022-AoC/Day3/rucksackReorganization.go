package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func calVal(c rune) int {
    ascii := int(c)
    if unicode.IsUpper(c) {
        return ascii - 38
    } else {
        return ascii - 96
    }
}

func compareCompartment(c1, c2, c3 string) int{
    sum := 0
    set := map[rune]int{} 
    for _, c := range c1 {
        set[c] = 1
    }

    for _, s := range c2 {
        if _, found := set[s]; found && set[s] == 1 {
            set[s] += 1
        }
    }
    
    for _, s := range c3 {
        if _, found := set[s]; found && set[s] == 2 {
            set[s] += 1
            sum += calVal(s)
        }
    }
    return sum
}

func getLine(r *bufio.Reader) (string, error) {
    line, err := r.ReadString('\n')
	if err != nil {
	    return "", err
	}
	return strings.TrimSuffix(line, "\n"), nil

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	total := 0
	for {
	    line1, err := getLine(reader)
	    if err != nil {
	        fmt.Println(total)
	        os.Exit(0)
	        
	    }
	    line2, _ := getLine(reader)
	    line3, _ := getLine(reader)
	    total += compareCompartment(line1, line2, line3)       
	}
}
