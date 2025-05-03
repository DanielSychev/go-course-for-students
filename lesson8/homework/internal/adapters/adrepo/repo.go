package adrepo

import (
	"errors"
	"homework8/internal/ads"
	"homework8/internal/app"
	"sync"
)

var ErrNotAuthor = errors.New("not author")
var ErrValidate = errors.New("validation error")

type Repo struct {
	index int64
	a     []ads.Ad
	mu    *sync.Mutex
}

func validate(Title string, Text string) bool {
	return Title != "" && len(Title) < 100 && Text != "" && len(Text) < 500
}

func (r *Repo) Create(Title string, Text string, UserID int64) (*ads.Ad, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if !validate(Title, Text) {
		return nil, ErrValidate
	}
	r.a = append(r.a, ads.Ad{
		ID:        r.index,
		Title:     Title,
		Text:      Text,
		AuthorID:  UserID,
		Published: false,
	})
	r.index++
	return &r.a[r.index-1], nil
}

func (r *Repo) UpdatePublished(ID int64, UserID int64, Published bool) (*ads.Ad, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.a[ID].AuthorID != UserID {
		return nil, ErrNotAuthor
	}
	r.a[ID].Published = Published
	return &r.a[ID], nil
}

func (r *Repo) UpdateTextAndTitle(ID int64, UserID int64, Title string, Text string) (*ads.Ad, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.a[ID].AuthorID != UserID {
		return nil, ErrNotAuthor
	}
	if !validate(Title, Text) {
		return nil, ErrValidate
	}
	r.a[ID].Text = Text
	r.a[ID].Title = Title
	return &r.a[ID], nil
}

func (r *Repo) GetList() ([]ads.Ad, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var res = make([]ads.Ad, 0)
	for _, elem := range r.a {
		if elem.Published {
			res = append(res, elem)
		}
	}
	return res, nil
}

func New() app.Repository {
	return &Repo{index: 0, a: []ads.Ad{}, mu: new(sync.Mutex)}
}
