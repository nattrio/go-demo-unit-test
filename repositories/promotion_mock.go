package repositories

import "github.com/stretchr/testify/mock"

type PromotionRepositoryMock struct {
	mock.Mock
}

// GetPromotion is a mock function for PromotionRepository.GetPromotion
func NewPromotionRepositoryMock() *PromotionRepositoryMock {
	return &PromotionRepositoryMock{}
}

func (m *PromotionRepositoryMock) GetPromotion() (Promotion, error) {
	args := m.Called()

	return args.Get(0).(Promotion), args.Error(1) // type assertion to Promotion
}
