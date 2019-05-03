package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
	"github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
)

func getArticles() {
	baseURL := "https://newsapi.org/v2/top-headlines"
	apiKeyPath := "./secret/apiKey"
	f, _ := os.Open(apiKeyPath)
	defer f.Close()

	buf, _ := ioutil.ReadAll(f)
	apiKey := string(buf)
	country := "jp"
	page := "1"
	pageSize := "100"
	categories := [7]string{"business", "entertainment", "general", "health", "science", "sports", "technology"}
	// categories := [1]string{"science"}
	for _, category := range categories {
		req, _ := http.NewRequest("GET", baseURL, nil)
		query := req.URL.Query()
		query.Add("country", country)
		query.Add("apiKey", apiKey)
		query.Add("pageSize", pageSize)
		query.Add("page", page)
		query.Add("category", category)
		req.URL.RawQuery = query.Encode()
		client := new(http.Client)
		res, _ := client.Do(req)
		defer res.Body.Close()

		byteArray, _ := ioutil.ReadAll(res.Body)

		results := json2struct(category, byteArray)
		insertDB(results)
	}
}

func insertDB(articles []*Article) {
	// db名
	session, _ := mgo.Dial("mongodb://localhost/user_modeling")
	defer session.Close()
	// db名
	db := session.DB("user_modeling")

	// collection名
	collection := db.C("articles")
	for _, article := range articles {
		if err := collection.Insert(article); err != nil {
			log.Fatal(err)
		}
	}

	var result Article
	err := collection.Find(bson.M{"category": "business"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.title, result.title)
}

func json2struct(category string, json []byte) []*Article {
	total := gjson.Get(string(json), "articles.#").Int()
	titles := gjson.Get(string(json), "articles.#.title").Array()
	descripitons := gjson.Get(string(json), "articles.#.description").Array()
	var articles []*Article
	fmt.Println(total)
	for i := 0; i < int(total); i++ {
		article := &Article{
			ID: bson.NewObjectId(),
			category: category,
			title: titles[i].String(),
			description: descripitons[i].String(),
		}
		articles = append(articles, article)
	}
	return articles
} 