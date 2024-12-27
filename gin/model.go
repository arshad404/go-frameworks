package main

type blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type getRequestBody struct {
	ID int `json:"id"`
}
