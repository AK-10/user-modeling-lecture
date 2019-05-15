package main

import (
    "fmt"
    "flag"
    "os"
    "strconv"
    "log"
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
        userID, _ := strconv.Atoi(flag.Arg(1))
        user, err := getUser(int64(userID))
        if err != nil {
            fmt.Println("user not found. invalid user ID.")
            fmt.Println("user ID is range 151 to 200")
            log.Fatal(err)
        }
        fmt.Println("========= user infomation =========")
        user.toString()
        fmt.Println("===================================")

        // others, err := getUsers()
        // if err != nil {
        //     log.Fatal(err)
        // }

    default:
        fmt.Println("Error Invalid argument")
        os.Exit(1)
    }
}
