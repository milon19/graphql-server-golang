package domain

type Author struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
