package usecase

import (
	"report-service/internal/model"
	"report-service/internal/repository"
	"strings"
)

type Service interface {
	CountByType(req model.CountByTypeRequest) (model.CountByTypeResponse, error)
	ListByColor(req model.ListByColorRequest) (model.ListByColorResponse, error)
}

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) CountByType(req model.CountByTypeRequest) (model.CountByTypeResponse, error) {
	lines, err := s.repo.Load()
	if err != nil {
		return model.CountByTypeResponse{}, err
	}

	count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if fields[2] == req.Tipe.String() {
			count++
		}
	}

	return model.CountByTypeResponse{JumlahKendaraan: count}, nil
}

func (s *serviceImpl) ListByColor(req model.ListByColorRequest) (model.ListByColorResponse, error) {
	lines, err := s.repo.Load()
	if err != nil {
		return model.ListByColorResponse{}, err
	}

	var plates []string
	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if fields[1] == req.Warna.String() {
			plates = append(plates, fields[0])
		}
	}

	return model.ListByColorResponse{PlatNomor: plates}, nil
}
