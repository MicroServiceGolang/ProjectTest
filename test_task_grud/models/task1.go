package models

type OpenApi struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
type Data struct {
	Data []OpenApi `json:"data"`
}
