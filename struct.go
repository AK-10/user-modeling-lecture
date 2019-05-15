package main

type User struct {
	ID   int64
	name string
	age  int64

	articles []*Article
	points map[string]int64
}


type Article struct {
	ID          int64 
	category    string        
	title       string        
	description string      
}

type ArticleResume struct {
	ID int64
	userID int64
	articleID int64
}

// {
// "business": 0,
// "entertainment": 0,
// "general": 0,
// "health": 0,
// "science": 0,
// "sports": 0,
// "technology": 0
// } 