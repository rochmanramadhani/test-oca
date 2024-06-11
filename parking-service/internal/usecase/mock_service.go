package usecase

import (
	"github.com/stretchr/testify/mock"
	"parking-service/internal/model"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) RegisterVehicle(req model.RegisterVehicleRequest) (model.RegisterVehicleResponse, error) {
	args := m.Called(req)
	return args.Get(0).(model.RegisterVehicleResponse), args.Error(1)
}

func (m *MockService) ExitVehicle(req model.ExitVehicleRequest) (model.ExitVehicleResponse, error) {
	args := m.Called(req)
	return args.Get(0).(model.ExitVehicleResponse), args.Error(1)
}
