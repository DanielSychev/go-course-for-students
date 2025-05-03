package httpgin

import (
	"github.com/gin-gonic/gin"
	"homework8/internal/ads"
)

type createAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

type adResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	AuthorID    int64  `json:"author_id"`
	Published   bool   `json:"published"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

type changeAdStatusRequest struct {
	Published bool  `json:"published"`
	UserID    int64 `json:"user_id"`
}

type updateAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

func AdSuccessResponse(ad *ads.Ad) gin.H {
	return gin.H{
		"data": adResponse{
			ID:          ad.ID,
			Title:       ad.Title,
			Text:        ad.Text,
			AuthorID:    ad.AuthorID,
			Published:   ad.Published,
			DateCreated: ad.DateCreated,
			DateUpdated: ad.DateUpdated,
		},
		"error": nil,
	}
}

func AdListSuccessResponse(ad []ads.Ad) gin.H {
	return gin.H{
		"data":  ad,
		"error": nil,
	}
}

func AdErrorResponse(err error) gin.H {
	return gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}
