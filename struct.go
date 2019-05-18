package main

import (
	"fmt"
)

type User struct {
	ID   int64
	name string
	age  int64

	articles []*Article
	points map[string]int64
	similarity float64
}

func (u *User) toString() {
	fmt.Println("id:", u.ID)
	fmt.Println("name:", u.name)
	fmt.Println("age:", u.age)
	// fmt.Println("resumes:", u.articles)
	categories := [7]string{"business", "entertainment", "general", "health", "science", "sports", "technology"}
	fmt.Println("------- points --------")
	for _, v := range categories {
		fmt.Println(v, ": ", u.points[v])
	}
	fmt.Println("-----------------------")
}


type Article struct {
	ID          int64
	category    string
	title       string
	description string
}


func (a *Article) toStringWithoutDescription() {
	fmt.Println("id: ", a.ID)
	fmt.Println("category: ", a.category)
	fmt.Println("title: ", a.title)
}

type ArticleResume struct {
	ID int64
	userID int64
	articleID int64
}
