package rest

import (
	"lib"
	"net/http"
	"report-service/internal/model"
	"report-service/internal/usecase"
)

type Handler interface {
	HealthCheck(http.ResponseWriter, *http.Request)
	CountByType(http.ResponseWriter, *http.Request)
	ListByColor(http.ResponseWriter, *http.Request)
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

func (h handlerImpl) CountByType(w http.ResponseWriter, r *http.Request) {
	request := model.CountByTypeRequest{
		Tipe: lib.CarType(r.URL.Query().Get("q")),
	}
	err := request.Validate()
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	response, err := h.service.CountByType(request)
	lib.WriteResponse(w, err, response)
}

func (h handlerImpl) ListByColor(w http.ResponseWriter, r *http.Request) {
	request := model.ListByColorRequest{
		Warna: lib.Color(r.URL.Query().Get("q")),
	}
	err := request.Validate()
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	response, err := h.service.ListByColor(request)
	lib.WriteResponse(w, err, response)
}
