package main

import (
     "fmt"
     "strings"
)


func main() {
	var i string
        fmt.Scanln(&i)
        var init int
        split1 := strings.Split(i, ";")
        for i := 0; i < len(split1); i++ {
            if split1[i] == "SELECT" {
               init = 1
               break;
            } else if split1[i] == "CREATE" {
               init = 2
               break;
            }
        }
        
        file := `1{Table
1}1[1d;id, 2d;name
--1}1d]1-2
--1}2d]Sergio-Ana`
        
        if init == 1 {
           split2 := strings.Split(file, "{")
           for a := 0; a < len(split2); a++ {
               split3 := strings.Split(split2[a], "[")
               if strings.Contains(split3[0], split1[3]) == true {
                  if split1[1] == "*" {
                     fmt.Println(split3[1])
                     break;
                  } else {
                     var bk bool = false
                     split4 := strings.Split(split3[1], "--")
                     s5 := strings.Split(split4[0], ",")
                     for b := 0; b < len(s5); b++ {
                         if strings.Contains(s5[b], split1[1]) {
                            s6 := strings.Split(s5[b], ";")
                            for c := 1; c < len(split4); c++ {
                                six := strings.TrimSpace(s6[0])
                                if strings.Contains(split4[c], six) == true {
                                   var1 := strings.Split(split4[c], "]")
                                   fmt.Println(var1[1])
                                   bk = true
                                }
                            }
                         }
                         if bk == true {
                            break;
                         }
                     }
                  }
               }
           }
        } 
}





