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
	sp := lib.CreateRootSpan(r, "HealthCheck")
	defer sp.Finish()

	lib.WriteResponse(sp, w, nil, "OK")
}

func (h handlerImpl) CountByType(w http.ResponseWriter, r *http.Request) {
	sp := lib.CreateRootSpan(r, "CountByType")
	defer sp.Finish()

	request := model.CountByTypeRequest{
		Tipe: lib.CarType(r.URL.Query().Get("q")),
	}
	lib.LogRequest(sp, request)

	err := request.Validate()
	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}

	response, err := h.service.CountByType(request)
	lib.WriteResponse(sp, w, err, response)
}

func (h handlerImpl) ListByColor(w http.ResponseWriter, r *http.Request) {
	sp := lib.CreateRootSpan(r, "ListByColor")
	defer sp.Finish()

	request := model.ListByColorRequest{
		Warna: lib.Color(r.URL.Query().Get("q")),
	}
	lib.LogRequest(sp, request)

	err := request.Validate()
	if err != nil {
		lib.WriteResponse(sp, w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	response, err := h.service.ListByColor(request)
	lib.WriteResponse(sp, w, err, response)
}
