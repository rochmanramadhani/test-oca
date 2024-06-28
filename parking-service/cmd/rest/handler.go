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
	sp := lib.CreateRootSpan(r, "HealthCheck")
	defer sp.Finish()

	lib.WriteResponse(sp, w, nil, "OK")
}

func (h handlerImpl) RegisterVehicle(w http.ResponseWriter, r *http.Request) {
	sp := lib.CreateRootSpan(r, "RegisterVehicle")
	defer sp.Finish()

	var request model.RegisterVehicleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	lib.LogRequest(sp, request)

	err = request.Validate()
	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	response, err := h.service.RegisterVehicle(request)
	lib.WriteResponse(sp, w, err, response)
}

func (h handlerImpl) ExitVehicle(w http.ResponseWriter, r *http.Request) {
	sp := lib.CreateRootSpan(r, "ExitVehicle")
	defer sp.Finish()

	var request model.ExitVehicleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	lib.LogRequest(sp, request)

	err = request.Validate()
	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	response, err := h.service.ExitVehicle(request)
	lib.WriteResponse(sp, w, err, response)
}
