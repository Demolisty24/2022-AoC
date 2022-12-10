package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	total := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(0)
		}
		return strings.TrimSuffix(line, "\n"), nil
	

	    if err != nil {
	        fmt.Println(total)
	        os.Exit(0)
	        
	    }
	    line2, _ := getLine(reader)
	    line3, _ := getLine(reader)
	    total += compareCompartment(line1, line2, line3)       
	}
}
