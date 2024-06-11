package repository

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockRepository) Load() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockRepository) Delete(plateNumber string) error {
	args := m.Called(plateNumber)
	return args.Error(0)
}
