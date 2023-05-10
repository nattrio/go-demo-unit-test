package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	// Create an instance of our test struct
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("John", 20, nil)
	c.On("GetCustomer", 2).Return("Mary", 30, errors.New("not found"))

	name, age, err := c.GetCustomer(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, age)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

// GetCustomer is a mock function for CustomerRepository.GetCustomer
func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2) // 0 is the first argument, 1 is the second, and so on...
}
