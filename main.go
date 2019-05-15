package main

import (
    "fmt"
    "flag"
    "os"
)


func main() {
    flag.Parse()
    arg := flag.Arg(0)
    switch arg {
    case "get-articles": // 済
        getArticles()
    case "make-user":
        makeUsers() // 済
    case "exec":
    default:
        fmt.Println("Error Invalid argument")
        os.Exit(1)
    }
    fmt.Println("Hello, User Modeling!")
}
