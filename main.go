package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    fmt.Printf("Enter your name please: ")
    var name string
    fmt.Scanf("%s", &name)
    fmt.Println("Rolling the dice for ", name)

    for true {
        rand.Seed(time.Now().UnixNano())
        random := 1 + rand.Intn(5)

        fmt.Println("You got ", random)
        if random == 6 {
            fmt.Println("Rolling again...")
        } else {
            break
        }
    }
}
