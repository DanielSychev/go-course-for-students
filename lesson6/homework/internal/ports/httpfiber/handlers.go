package httpfiber

import (
	"errors"
	"homework6/internal/adapters/adrepo"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"homework6/internal/app"
)

// Метод для создания объявления (ad)
func createAd(a app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody createAdRequest
		err := c.BodyParser(&reqBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(AdErrorResponse(err))
		}

		//TODO: вызов логики, например, CreateAd(c.Context(), reqBody.Title, reqBody.Text, reqBody.UserID)
		// TODO: метод должен возвращать AdSuccessResponse или ошибку.

		ad, err := a.CreateAd(c.Context(), reqBody.Title, reqBody.Text, reqBody.UserID)

		if err != nil {
			if errors.Is(err, adrepo.ErrNotAuthor) {
				c.Status(http.StatusForbidden)
			} else if errors.Is(err, adrepo.ErrValidate) {
				c.Status(http.StatusBadRequest)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			return c.JSON(AdErrorResponse(err))
		}

		return c.JSON(AdSuccessResponse(ad))
	}
}

// Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
func changeAdStatus(a app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody changeAdStatusRequest
		if err := c.BodyParser(&reqBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(AdErrorResponse(err))
		}

		adID, err := c.ParamsInt("ad_id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(AdErrorResponse(err))
		}

		// TODO: вызов логики ChangeAdStatus(c.Context(), int64(adID), reqBody.UserID, reqBody.Published)
		// TODO: метод должен возвращать AdSuccessResponse или ошибку.

		ad, err := a.ChangeAdStatus(c.Context(), int64(adID), reqBody.UserID, reqBody.Published)

		if err != nil {
			if errors.Is(err, adrepo.ErrNotAuthor) {
				c.Status(http.StatusForbidden)
			} else if errors.Is(err, adrepo.ErrValidate) {
				c.Status(http.StatusBadRequest)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			return c.JSON(AdErrorResponse(err))
		}

		return c.JSON(AdSuccessResponse(ad))
	}
}

// Метод для обновления текста(Text) или заголовка(Title) объявления
func updateAd(a app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody updateAdRequest
		if err := c.BodyParser(&reqBody); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(AdErrorResponse(err))
		}

		adID, err := c.ParamsInt("ad_id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(AdErrorResponse(err))
		}

		ad, err := a.UpdateAd(c.Context(), int64(adID), reqBody.UserID, reqBody.Title, reqBody.Text)

		// TODO: вызов логики, например, UpdateAd(c.Context(), int64(adID), reqBody.UserID, reqBody.Title, reqBody.Text)
		// TODO: метод должен возвращать AdSuccessResponse или ошибку.

		if err != nil {
			if errors.Is(err, adrepo.ErrNotAuthor) {
				c.Status(http.StatusForbidden)
			} else if errors.Is(err, adrepo.ErrValidate) {
				c.Status(http.StatusBadRequest)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			return c.JSON(AdErrorResponse(err))
		}

		return c.JSON(AdSuccessResponse(ad))
	}
}
