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

        others, err := getUsers()
        if err != nil {
            log.Fatal(err)
        }
        
        cosineSimilarity(user, others)
        
        fmt.Println("============ similars =============")
        deleteMyself(user, others)
        similars := getSimilars(0.7, others)
        sortBySimilarity(similars)
        for _, sim := range similars {
            fmt.Println("id:", sim.ID)
            fmt.Println("name: ", sim.name)
            fmt.Println("similarity: ", sim.similarity)
        	categories := [7]string{"business", "entertainment", "general", "health", "science", "sports", "technology"}
            fmt.Println("------- points --------")
            for _, v := range categories {
                fmt.Println(v, ": ", sim.points[v])
            }
        }
    	fmt.Println("-----------------------") 
        fmt.Println("===================================")
        
        fmt.Println("====== recommended articles =======")
        recommended := getRecommendedArticles(user, similars)
        for _, article := range recommended {
            article.toStringWithoutDescription()
            fmt.Println("")
        }
        fmt.Println("===================================")

        fmt.Println("======= recommended counts ========")
        counter := map[string]int{
            "business": 0,
            "entertainment": 0,
            "general": 0,
            "health": 0,
            "science": 0,
            "sports": 0,
            "technology": 0,
        }
        for _, article := range recommended {
            counter[article.category]++
        }
        fmt.Println("recommended counter: ", counter)
        fmt.Println("===================================")

        // pointsを正規化: 単位ベクトルに変換
        // 内積を取ればコサイン類似度が出る
    default:
        fmt.Println("Error Invalid argument")
        os.Exit(1)
    }
}

