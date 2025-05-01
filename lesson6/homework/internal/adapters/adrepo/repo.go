package adrepo

import (
	"errors"
	"homework6/internal/ads"
	"homework6/internal/app"
)

var ErrNotAuthor = errors.New("not author")
var ErrValidate = errors.New("validation error")

type Repo struct {
	index int64
	a     []ads.Ad
}

func validate(Title string, Text string) bool {
	return Title != "" && len(Title) < 100 && Text != "" && len(Text) < 500
}

func (r *Repo) Create(Title string, Text string, UserID int64) (*ads.Ad, error) { // TODO: Сделать валидацию
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
	if r.a[ID].AuthorID != UserID {
		return nil, ErrNotAuthor
	}
	r.a[ID].Published = Published
	return &r.a[ID], nil
}

func (r *Repo) UpdateTextAndTitle(ID int64, UserID int64, Title string, Text string) (*ads.Ad, error) { // TODO: Сделать валидацию
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

func New() app.Repository {
	return &Repo{index: 0, a: []ads.Ad{}} // TODO: реализовать
}
