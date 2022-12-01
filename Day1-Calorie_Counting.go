package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
    )

func main() {
    var calories = make([]float64,0)
    var current float64
    
    reader := bufio.NewReader(os.Stdin)
    
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            sort.Float64s(calories)
            calories = calories[len(calories)-3:]
            os.Exit(0)
        }
        
        line = strings.TrimSpace(line)
        if len(line) == 0 {
            calories = append(calories, current)
            current = 0
        } else {
            calorie, _ := strconv.ParseFloat(line, 64)
            current += calorie
        }
    
    }
}
