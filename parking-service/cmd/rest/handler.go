package rest

import (
	"encoding/json"
	"lib"
	"net/http"
	"parking-service/internal/model"
	"parking-service/internal/usecase"
)

type Handler interface {
	HealthCheck(http.ResponseWriter, *http.Request)
	RegisterVehicle(http.ResponseWriter, *http.Request)
	ExitVehicle(http.ResponseWriter, *http.Request)
}

type handlerImpl struct {
	service usecase.Service
}

func NewHandler(service usecase.Service) *handlerImpl {
	return &handlerImpl{service: service}
}

func (h handlerImpl) HealthCheck(w http.ResponseWriter, r *http.Request) {
	lib.WriteResponse(w, nil, "OK")
}

func (h handlerImpl) RegisterVehicle(w http.ResponseWriter, r *http.Request) {
	var request model.RegisterVehicleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	err = request.Validate()
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	response, err := h.service.RegisterVehicle(request)
	lib.WriteResponse(w, err, response)
}

func (h handlerImpl) ExitVehicle(w http.ResponseWriter, r *http.Request) {
	var request model.ExitVehicleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	err = request.Validate()
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	response, err := h.service.ExitVehicle(request)
	lib.WriteResponse(w, err, response)
}
