package handlers_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-demo-unit-test/handlers"
	"github.com/nattrio/go-demo-unit-test/services"
	"github.com/stretchr/testify/assert"
)

func TestPromotionHandlerCalculateDiscount(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		// Arange
		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			// If status code is OK, then we can read the response body
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, "strconv.Itoa(expected)", string(body))
		}

	})
}
