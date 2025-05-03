package ads

type Ad struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	AuthorID    int64  `json:"author_id"`
	Published   bool   `json:"published"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

type AdFilter struct {
	Pub   bool
	Auth  int64
	Title string
}
