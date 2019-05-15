package main

import (
	"errors"
	"fmt"
	"database/sql"
	"encoding/csv"
    "log"
	"os"
	"io"
	"strconv"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

// 全員が20個の閲覧履歴を持っているものとする
// 完全ランダム
// 1つのカテゴリ
// 2つのカテゴリ
// 3つのカテゴリ
// 4つのカテゴリ

func makeUser(name string, age int64, option int) (*User, error) {
	user := User{name: name, age: age}
	 
	dsName := getDBName()
	db, err := sql.Open("mysql", dsName)
	if err != nil {
		panic(err.Error())
	}
	ins, err := db.Prepare("INSERT INTO users(name, age) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()
	result, err := ins.Exec(user.name, user.age)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("cannot insert userdata")
	} else {
		userID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		user.ID = userID
		switch option {
		case 0:
			rows, err := db.Query("SELECT id FROM articles ORDER BY RAND() LIMIT 20")
			if err != nil {
				log.Fatal(err)
			}
			var articleIDs [20]int64
			index := 0
			for rows.Next() {
				if err := rows.Scan(&articleIDs[index]); err != nil {
					log.Fatal(err)
				}
				index++
			}
			for i := 0; i < 20; i++ {
				if err := insertResume(userID, articleIDs[i], db); err != nil {
					log.Fatal(err)
				}
			}
		case 1, 2, 3:
			makeResume(user.ID, option, db)
		default:
			return nil, errors.New("invalid option number")
		}
	}
	return &user, nil
}



func insertResume(userID int64, articleID int64, db *sql.DB) error {
	ins, err := db.Prepare("INSERT INTO article_resumes(user_id, article_id) VALUES(?,?)")
	if err != nil {
		return errors.New("invalid db")
		
	}
	defer ins.Close()
	_, err = ins.Exec(userID, articleID)

	if err != nil {
		return errors.New("could not insert data")
	}
	return nil
}

func makeResume(userID int64, num int, db *sql.DB) {
	categories := [7]string{"business", "entertainment", "general", "health", "science", "sports", "technology"}	
	rand.Seed(userID)
	indices := rand.Perm(7)[:num]
	var selectedCategories []string
	for _, idx := range indices {
		selectedCategories = append(selectedCategories, categories[idx])
	}
	condition := ""
	for i:= 0; i < len(selectedCategories); i++ {
		condition += "category = \"" + selectedCategories[i] + "\""
		if i != len(selectedCategories) - 1 {
			condition += " OR "
		}
	}
	rows, err := db.Query("SELECT id FROM articles WHERE " + condition + " ORDER BY RAND() LIMIT 20")
	if err != nil {
		log.Fatal(err)
	}
	var articleIDs [20]int64
	index := 0
	for rows.Next() {
		if err := rows.Scan(&articleIDs[index]); err != nil {
			log.Fatal(err)
		}
		index++
	}
	for i := 0; i < 20; i++ {
		err := insertResume(userID, articleIDs[i], db)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func makeUsers() {
	// var users []*User
	file, err := os.Open("./users.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string
	for {
		line, err = reader.Read()

		if err == io.EOF {
			break
		} 
		if err != nil {
			log.Fatal(err)
		}
		age, _ := strconv.Atoi(line[1])
		_, err := makeUser(line[0], int64(age), age % 4)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func getUser(userID int64) (*User, error) {
	dsName := getDBName()
	db, err := sql.Open("mysql", dsName)
	println("x")
	if err != nil {
		return nil, err
	}
	
	points := map[string]int64{
		"business": 0,
		"entertainment": 0,
		"general": 0,
		"health": 0,
		"science": 0,
		"sports": 0,
		"technology": 0,
	}
	var user User
	user.points = points
	if err := db.QueryRow("SELECT * FROM users u WHERE u.id = ?", userID).Scan(&user.ID, &user.name, &user.age); err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT a.* FROM article_resumes ar INNER JOIN articles a ON ar.article_id = a.id WHERE ar.user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.category, &article.title, &article.description); err != nil {
			return nil, err
		}
		user.articles = append(user.articles, &article)
		user.points[article.category]++
	}
	return &user, nil
}

func getUsers() ([]*User, error) {
	dsName := getDBName()
	db, err := sql.Open("mysql", dsName)
	println("x")
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT id FROM users")
	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var userID int64 
		if rows.Scan(&userID); err != nil {
			return nil, err
		}
		user, err := getUser(userID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}


func (u *User) toString() {
	fmt.Println("id:", u.ID)
	fmt.Println("name:", u.name)
	fmt.Println("age:", u.age)
	// fmt.Println("resumes:", u.articles)
	fmt.Println("points:", u.points)
}
