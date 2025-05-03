package httpgin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"homework8/internal/adapters/adrepo"
	"homework8/internal/app"
	"log"
	"net/http"
	"strconv"
)

func CreateHandle(c *gin.Context, a app.App) {
	var adReq createAdRequest

	if err := c.BindJSON(&adReq); err != nil {
		c.JSON(http.StatusBadRequest, AdErrorResponse(err))
		return
	}
	jsonData, err := json.Marshal(adReq)
	log.Println(string(jsonData))
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
	jsonData, err := json.Marshal(adReq)
	log.Println(string(jsonData))
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
	jsonData, err := json.Marshal(adReq)
	log.Println(string(jsonData))
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
	adResp, err := a.GetList(c)

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
