package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Body of article 1"},
	{ID: 2, Title: "Go is Cool", Content: "Now I can do full stack UI/API/DB in GO why go anywhere else?"},
}

func getAllArticles() []article {
	return articleList
}
