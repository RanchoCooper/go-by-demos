package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Response struct {
	Page       int32     `json:"page"`
	PerPage    int32     `json:"per_page"`
	Total      int32     `json:"total"`
	TotalPages int32     `json:"total_pages"`
	Data       []Article `json:"data"`
}

type Article struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Author      string `json:"author"`
	NumComments int32  `json:"num_comments"`
	StoryId     int32  `json:"story_id"`
	StoryTitle  string `json:"story_title"`
	StoryUrl    string `json:"story_url"`
	ParentId    int32  `json:"parent_id"`
	CreatedAt   int32  `json:"created_at"`
}

func topArticles(limit int32) []string {
	results := make([]string, 0)

	for i := 1; i <= int(limit); i++ {
		resp, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/articles?page=%d", i))
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		response := Response{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		maxComments := 0
		topArticle := Article{}
		for _, article := range response.Data {
			if article.NumComments > int32(maxComments) {
				if article.Title != "" || article.StoryUrl != "" {
					maxComments = int(article.NumComments)
					topArticle = article
				}
			}
		}
		if topArticle.Title != "" {
			results = append(results, topArticle.Title)
		} else if topArticle.StoryTitle != "" {
			results = append(results, topArticle.StoryTitle)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})

	return results
}

func main() {

}
