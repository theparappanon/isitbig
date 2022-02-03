package main

import "fmt"

func main() {
    
    number := 0

    fmt.Println("Type a number")
    fmt.Scanf("%d", &number)

    if number > 100 {
        fmt.Println("This number is big.")
       } else {
        fmt.Println("This number isn't big.")
       }
    }
