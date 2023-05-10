package services

import (
	"github.com/nattrio/go-demo-unit-test/repositories"
)

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionService struct {
	// PromoRepo is an interface that is implemented by the PromotionRepository struct
	PromoRepo repositories.PromotionRepository
}

// NewPromotionService returns a new instance of PromotionService
func NewPromotionService(promoRepo repositories.PromotionRepository) PromotionService {
	return promotionService{PromoRepo: promoRepo}
}

// CalculateDiscount calculates the discount based on the amount
func (s promotionService) CalculateDiscount(amount int) (int, error) {
	if amount <= 0 {
		return 0, ErrZeroAmount
	}

	// GetPromotion is a method of the PromotionRepository interface
	promotion, err := s.PromoRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	// If the amount is greater than the minimum purchase amount, apply the discount
	if amount >= promotion.PurchaseMin {
		return amount - (promotion.DiscountPercent * amount / 100), nil
	}

	// Otherwise, return the amount without discount
	return amount, nil
}
