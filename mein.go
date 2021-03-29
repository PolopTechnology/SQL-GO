package main

import (
     "fmt"
     "strings"
)


func main() {
	var i string
        fmt.Scanln(&i)
        b := [4]string{"SELECT", "*", "FROM", "Table"}
	syntax := strings.Split(i, ";")
        var points int
        for x := 0; x < len(syntax); x++ {
            for j := x; j < x + 1; j++ {
                if syntax[x] == b[j] {
                   points += 1
                }
            }
        } 
        if points >= 4 {
           fmt.Println("id | name")
        } else if points > 0 && points < 4 {
           fmt.Println("SYNTAX ERROR! INCOMPLETE SENTENCE!")
        }
}





