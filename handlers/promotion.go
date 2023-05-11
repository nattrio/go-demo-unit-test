package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-demo-unit-test/services"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promoService services.PromotionService
}

// NewPromotionHandler is a constructor for promotionHandler
func NewPromotionHandler(promoService services.PromotionService) PromotionHandler {
	return promotionHandler{promoService: promoService} // return a new instance of promotionHandler
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {

	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	discount, err := h.promoService.CalculateDiscount(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendString(strconv.Itoa(discount))
}
