package usecase

import (
	"github.com/stretchr/testify/mock"
	"report-service/internal/model"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) CountByType(req model.CountByTypeRequest) (model.CountByTypeResponse, error) {
	args := m.Called(req)
	return args.Get(0).(model.CountByTypeResponse), args.Error(1)
}

func (m *MockService) ListByColor(req model.ListByColorRequest) (model.ListByColorResponse, error) {
	args := m.Called(req)
	return args.Get(0).(model.ListByColorResponse), args.Error(1)
}
