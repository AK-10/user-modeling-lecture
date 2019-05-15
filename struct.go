package main

type User struct {
	ID   int64
	name string
	age  int64
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