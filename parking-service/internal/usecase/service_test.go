package usecase_test

import (
	"parking-service/internal/model"
	"parking-service/internal/repository"
	"parking-service/internal/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_RegisterVehicle(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	service := usecase.NewService(mockRepo)

	req := model.RegisterVehicleRequest{PlatNomor: "ABC123", Warna: "Red", Tipe: "SUV"}
	mockRepo.On("Save", mock.Anything).Return(nil)

	resp, err := service.RegisterVehicle(req)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, req.PlatNomor, resp.PlatNomor)
	assert.Equal(t, "A1", resp.ParkingLot)

	tanggalMasuk, err := time.Parse("2006-01-02 15:04", resp.TanggalMasuk)
	assert.NoError(t, err)
	loc, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	tanggalMasuk = tanggalMasuk.In(loc)
	now := time.Now().In(loc)
	assert.True(t, now.Sub(tanggalMasuk) < time.Second)
}

func TestService_ExitVehicle(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	service := usecase.NewService(mockRepo)

	req := model.ExitVehicleRequest{PlatNomor: "ABC123"}
	expectedLoadData := []string{"ABC123,Red,SUV,A1,2024-06-11 10:00"}
	mockRepo.On("Load").Return(expectedLoadData, nil)
	mockRepo.On("Delete", "ABC123").Return(nil)

	resp, err := service.ExitVehicle(req)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, req.PlatNomor, resp.PlatNomor)
	assert.Equal(t, "2024-06-11 10:00", resp.TanggalMasuk)

	tanggalKeluar, err := time.Parse("2006-01-02 15:04", resp.TanggalKeluar)
	assert.NoError(t, err)
	loc, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	tanggalKeluar = tanggalKeluar.In(loc)
	now := time.Now().In(loc)
	assert.True(t, now.Sub(tanggalKeluar) < time.Second)
	assert.Equal(t, 25000, resp.JumlahBayar)
}
