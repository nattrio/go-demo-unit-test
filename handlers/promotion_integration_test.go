//go:build integration

package handlers_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-demo-unit-test/handlers"
	"github.com/nattrio/go-demo-unit-test/repositories"
	"github.com/nattrio/go-demo-unit-test/services"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegrationService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		amount := 100
		expected := 80

		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)
		promoHandler := handlers.NewPromotionHandler(promoService)

		// http://localhost:8000/calculatediscount?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			// If status code is OK, then we can read the response body
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})
}
