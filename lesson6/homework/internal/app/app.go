package app

import (
	"github.com/valyala/fasthttp"
	"homework6/internal/ads"
)

type App interface {
	// TODO: реализовать
	CreateAd(c *fasthttp.RequestCtx, Title string, Text string, UserID int64) (*ads.Ad, error)
	ChangeAdStatus(c *fasthttp.RequestCtx, ID int64, UserID int64, Published bool) (*ads.Ad, error)
	UpdateAd(c *fasthttp.RequestCtx, ID int64, UserID int64, Title string, Text string) (*ads.Ad, error)
}

type Repository interface {
	// TODO: реализовать
	Create(Title string, Text string, UserID int64) (*ads.Ad, error)
	UpdatePublished(ID int64, UserID int64, Published bool) (*ads.Ad, error)
	UpdateTextAndTitle(ID int64, UserID int64, Title string, Text string) (*ads.Ad, error)
}

type AppMethods struct {
	r Repository
}

func (apm *AppMethods) CreateAd(c *fasthttp.RequestCtx, Title string, Text string, UserID int64) (*ads.Ad, error) { // TODO: Обработать ошибки валидации
	ad, err := apm.r.Create(Title, Text, UserID)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (apm *AppMethods) ChangeAdStatus(c *fasthttp.RequestCtx, ID int64, UserID int64, Published bool) (*ads.Ad, error) {
	ad, err := apm.r.UpdatePublished(ID, UserID, Published)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (apm *AppMethods) UpdateAd(c *fasthttp.RequestCtx, ID int64, UserID int64, Title string, Text string) (*ads.Ad, error) {
	ad, err := apm.r.UpdateTextAndTitle(ID, UserID, Title, Text)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func NewApp(repo Repository) App {
	return &AppMethods{r: repo} // TODO: реализовать
}
