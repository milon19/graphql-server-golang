package domain

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Price  float32 `json:"price"`
	IsbnNo int     `json:"isbn_no"`
}
