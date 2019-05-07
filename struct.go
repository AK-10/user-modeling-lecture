package main

type User struct {
	name string
	age  int
}


type Article struct {
	ID          bson.ObjectId `bson:"_id"`
	category    string        `bson:"category"`
	title       string        `bson:"title"`
	description string        `bson:"description"`
}
