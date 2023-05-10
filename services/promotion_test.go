package services_test

import (
	"errors"
	"testing"

	"github.com/nattrio/go-demo-unit-test/repositories"
	"github.com/nattrio/go-demo-unit-test/services"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name        string
		purchaseMin int
		discount    int
		amount      int
		expected    int
	}

	cases := []testCase{
		{name: "Applied 100", purchaseMin: 100, discount: 20, amount: 100, expected: 80},
		{name: "Applied 200", purchaseMin: 100, discount: 20, amount: 200, expected: 160},
		{name: "Applied 300", purchaseMin: 100, discount: 20, amount: 300, expected: 240},
		{name: "Not applied 50", purchaseMin: 100, discount: 20, amount: 50, expected: 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Arange
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discount,
			}, nil)

			promoService := services.NewPromotionService(promoRepo)

			// Act
			discount, _ := promoService.CalculateDiscount(c.amount)
			expected := c.expected

			// Assert
			assert.Equal(t, expected, discount)

		})
	}

	t.Run("Purchase amount is zero", func(t *testing.T) {
		// Arange
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoService.CalculateDiscount(0)

		// Assert
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion") // Assert that GetPromotion is not called
	})

	t.Run("Repository error", func(t *testing.T) {
		// Arange
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New(""))

		promoService := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoService.CalculateDiscount(100)

		// Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})
}
