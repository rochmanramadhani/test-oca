package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"parking-service/cmd/rest"
	"parking-service/internal/model"
	"parking-service/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_RegisterVehicle(t *testing.T) {
	mockService := new(usecase.MockService)
	handler := rest.NewHandler(mockService)

	reqBody := model.RegisterVehicleRequest{PlatNomor: "ABC123", Warna: "Red", Tipe: "SUV"}
	expectedResponse := model.RegisterVehicleResponse{PlatNomor: "ABC123", ParkingLot: "A1", TanggalMasuk: "2024-06-11 10:00"}
	mockService.On("RegisterVehicle", reqBody).Return(expectedResponse, nil)

	reqBodyBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.RegisterVehicle(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestHandler_ExitVehicle(t *testing.T) {
	mockService := new(usecase.MockService)
	handler := rest.NewHandler(mockService)

	reqBody := model.ExitVehicleRequest{PlatNomor: "ABC123"}
	expectedResponse := model.ExitVehicleResponse{PlatNomor: "ABC123", TanggalMasuk: "2024-06-11 10:00", TanggalKeluar: "2024-06-11 12:00", JumlahBayar: 25000}
	mockService.On("ExitVehicle", reqBody).Return(expectedResponse, nil)

	reqBodyBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/exit", bytes.NewBuffer(reqBodyBytes))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ExitVehicle(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}
