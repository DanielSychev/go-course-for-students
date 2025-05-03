package app

import (
	"github.com/gin-gonic/gin"
	"homework8/internal/ads"
)

type App interface {
	CreateAd(c *gin.Context, Title string, Text string, UserID int64) (*ads.Ad, error)
	ChangeAdStatus(c *gin.Context, ID int64, UserID int64, Published bool) (*ads.Ad, error)
	UpdateAd(c *gin.Context, ID int64, UserID int64, Title string, Text string) (*ads.Ad, error)
	GetList(c *gin.Context, filter ads.AdFilter) ([]ads.Ad, error)
	GetByID(c *gin.Context, ID int64) (*ads.Ad, error)
}

type Repository interface {
	Create(Title string, Text string, UserID int64) (*ads.Ad, error)
	UpdatePublished(ID int64, UserID int64, Published bool) (*ads.Ad, error)
	UpdateTextAndTitle(ID int64, UserID int64, Title string, Text string) (*ads.Ad, error)
	GetList(filter ads.AdFilter) ([]ads.Ad, error)
	GetByID(ID int64) (*ads.Ad, error)
}

type AppMethods struct {
	r Repository
}

func (apm *AppMethods) CreateAd(c *gin.Context, Title string, Text string, UserID int64) (*ads.Ad, error) {
	ad, err := apm.r.Create(Title, Text, UserID)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (apm *AppMethods) ChangeAdStatus(c *gin.Context, ID int64, UserID int64, Published bool) (*ads.Ad, error) {
	ad, err := apm.r.UpdatePublished(ID, UserID, Published)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (apm *AppMethods) UpdateAd(c *gin.Context, ID int64, UserID int64, Title string, Text string) (*ads.Ad, error) {
	ad, err := apm.r.UpdateTextAndTitle(ID, UserID, Title, Text)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (apm *AppMethods) GetList(c *gin.Context, filter ads.AdFilter) ([]ads.Ad, error) {
	return apm.r.GetList(filter)
}

func (apm *AppMethods) GetByID(c *gin.Context, ID int64) (*ads.Ad, error) {
	return apm.r.GetByID(ID)
}

func NewApp(repo Repository) App {
	return &AppMethods{r: repo}
}
