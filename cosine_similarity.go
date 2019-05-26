package main

import (
	"math"
	"sort"
	"fmt"
)

func normalizePoints(user *User) [7]float64 {
	scalar := 0.0
	for _, v := range user.points {
		scalar += math.Pow(float64(v), 2)
	}
	scalar = math.Sqrt(scalar)
	var normedPoints [7]float64
	idx := 0
	categories := [7]string{"business", "entertainment", "general", "health", "science", "sports", "technology"}
	for i, category := range categories {
		// divide all points by scalar
		normedPoints[i] = float64(user.points[category]) / scalar
		idx++
	}

	return normedPoints
}

func cosineSimilarity(user *User, users []*User) {
	userNormedPoints := normalizePoints(user)
	for _, u := range users {
		normed := normalizePoints(u)
		similarity := 0.0
		for i := 0; i < 7; i++ {
			similarity += userNormedPoints[i] * normed[i]
		}
		u.similarity = similarity
	}
}

func sortBySimilarity(users []*User) {
	sort.Slice(users, func(i, j int) bool { return users[i].similarity > users[j].similarity })
}

func deleteMyself(user *User, users []*User) {
	for i, u := range users {
		if user.ID == u.ID {
			users = append(users[:i], users[i+1:]...)
		}
	}
}

func getFrequencies(user *User) []string {
	var frequencies []string
	for k, v := range user.points {
		if v > 8.0 {
			frequencies = append(frequencies, k)
		}
	}
	return frequencies
} 

func getRecommendedArticles(user *User, users []*User) []*Article {
	// sort by similairty
	sortBySimilarity(users)
	deleteMyself(user, users)
	var recommendedArticles []*Article
	
	frequencies := getFrequencies(user)
	fmt.Println(frequencies)
	// pick up have not read article yet
	for _, s := range users {
		for _, a := range s.articles {
			if !(containsCategory(a.category, frequencies) || contains(a, user.articles)) {
				recommendedArticles = append(recommendedArticles, a)
			}
		}
	}
	return recommendedArticles
}

func containsCategory(category string, categories []string) bool {
	for _, c := range categories {
		if c == category {
			return true
		}
	}
	return false	
}

func contains(article *Article, articles []*Article) bool {
	for _, a := range articles {
		if article.ID == a.ID {
			return true
		}
	}
	return false
}