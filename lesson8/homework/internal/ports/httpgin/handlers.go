package httpgin

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"homework8/internal/adapters/adrepo"
	"homework8/internal/ads"
	"homework8/internal/app"
	"net/http"
	"strconv"
)

func CreateHandle(c *gin.Context, a app.App) {
	var adReq createAdRequest

	if err := c.BindJSON(&adReq); err != nil {
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	//jsonData, err := json.Marshal(adReq)
	//log.Println(string(jsonData))
	adResp, err := a.CreateAd(c, adReq.Title, adReq.Text, adReq.UserID)

	if err != nil {
		if errors.Is(err, adrepo.ErrNotAuthor) {
			c.JSON(http.StatusForbidden, AdErrorResponse(err))
		} else if errors.Is(err, adrepo.ErrValidate) {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		} else {
			c.JSON(http.StatusInternalServerError, AdErrorResponse(err))
		}
		return
	}
	c.JSON(http.StatusOK, AdSuccessResponse(adResp))
}

func ChangeAdStatus(c *gin.Context, a app.App) {
	var adReq changeAdStatusRequest

	strId := c.Param("id")
	adId, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, AdErrorResponse(fmt.Errorf("id should be a number")))
		return
	}

	if err := c.BindJSON(&adReq); err != nil {
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	//jsonData, err := json.Marshal(adReq)
	//log.Println(string(jsonData))
	adResp, err := a.ChangeAdStatus(c, adId, adReq.UserID, adReq.Published)

	if err != nil {
		if errors.Is(err, adrepo.ErrNotAuthor) {
			c.JSON(http.StatusForbidden, AdErrorResponse(err))
		} else if errors.Is(err, adrepo.ErrValidate) {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		} else {
			c.JSON(http.StatusInternalServerError, AdErrorResponse(err))
		}
		return
	}
	c.JSON(http.StatusOK, AdSuccessResponse(adResp))
}

func UpdateAd(c *gin.Context, a app.App) {
	var adReq updateAdRequest

	strId := c.Param("id")
	adId, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, AdErrorResponse(fmt.Errorf("id should be a number")))
		return
	}

	if err := c.BindJSON(&adReq); err != nil {
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	//jsonData, err := json.Marshal(adReq)
	//log.Println(string(jsonData))
	adResp, err := a.UpdateAd(c, adId, adReq.UserID, adReq.Title, adReq.Text)

	if err != nil {
		if errors.Is(err, adrepo.ErrNotAuthor) {
			c.JSON(http.StatusForbidden, AdErrorResponse(err))
		} else if errors.Is(err, adrepo.ErrValidate) {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		} else {
			c.JSON(http.StatusInternalServerError, AdErrorResponse(err))
		}
		return
	}
	c.JSON(http.StatusOK, AdSuccessResponse(adResp))
}

func GetList(c *gin.Context, a app.App) {
	var err error
	filter := ads.AdFilter{}
	if filter.Pub, err = strconv.ParseBool(c.Query("pub")); err != nil {
		filter.Pub = true // default: Published = true
	}
	if filter.Auth, err = strconv.ParseInt(c.Query("auth"), 10, 64); err != nil {
		filter.Auth = -1
	}
	filter.Title = c.Query("title")

	adResp, err := a.GetList(c, filter)

	if err != nil {
		if errors.Is(err, adrepo.ErrNotAuthor) {
			c.JSON(http.StatusForbidden, AdErrorResponse(err))
		} else if errors.Is(err, adrepo.ErrValidate) {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		} else {
			c.JSON(http.StatusInternalServerError, AdErrorResponse(err))
		}
		return
	}
	c.JSON(http.StatusOK, AdListSuccessResponse(adResp))
}

func GetAdById(c *gin.Context, a app.App) {
	strId := c.Param("id")
	adId, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, AdErrorResponse(fmt.Errorf("id should be a number")))
		return
	}

	adResp, err := a.GetByID(c, adId)

	if err != nil {
		if errors.Is(err, adrepo.ErrNotCreated) {
			c.JSON(http.StatusBadRequest, AdErrorResponse(fmt.Errorf("ad with such id hasn't created yet")))
		} else {
			c.JSON(http.StatusInternalServerError, AdErrorResponse(err))
		}
		return
	}
	c.JSON(http.StatusOK, AdSuccessResponse(adResp))
}
