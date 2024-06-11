package usecase

import (
	"fmt"
	"lib"
	"parking-service/internal/model"
	"parking-service/internal/repository"
	"strings"
	"time"
)

type Service interface {
	RegisterVehicle(req model.RegisterVehicleRequest) (model.RegisterVehicleResponse, error)
	ExitVehicle(req model.ExitVehicleRequest) (model.ExitVehicleResponse, error)
}

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) RegisterVehicle(req model.RegisterVehicleRequest) (model.RegisterVehicleResponse, error) {
	lines, err := s.repo.Load()
	if err != nil {
		return model.RegisterVehicleResponse{}, err
	}

	if err := s.validateRegistration(req.PlatNomor, lines); err != nil {
		return model.RegisterVehicleResponse{}, err
	}

	if len(lines) == lib.Cfg.MaxParkingLot {
		return model.RegisterVehicleResponse{}, lib.NewErrForbidden("parking lot is full")
	}

	parkingLot := s.generateParkingLot(lines)
	tanggalMasuk := time.Now().Format("2006-01-02 15:04")
	data := fmt.Sprintf("%s,%s,%s,%s,%s", req.PlatNomor, req.Warna, req.Tipe, parkingLot, tanggalMasuk)

	if err := s.repo.Save(data); err != nil {
		return model.RegisterVehicleResponse{}, err
	}

	return model.RegisterVehicleResponse{
		PlatNomor:    req.PlatNomor,
		ParkingLot:   parkingLot,
		TanggalMasuk: tanggalMasuk,
	}, nil
}

func (s *serviceImpl) ExitVehicle(req model.ExitVehicleRequest) (model.ExitVehicleResponse, error) {
	lines, err := s.repo.Load()
	if err != nil {
		return model.ExitVehicleResponse{}, err
	}

	if err := s.validateExit(req.PlatNomor, lines); err != nil {
		return model.ExitVehicleResponse{}, err
	}

	var response model.ExitVehicleResponse

	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if fields[0] == req.PlatNomor {
			response.PlatNomor = fields[0]
			response.TanggalMasuk = fields[4]
			response.TanggalKeluar = time.Now().Format("2006-01-02 15:04")

			hours := s.calculateHoursParked(response.TanggalMasuk, response.TanggalKeluar)
			response.JumlahBayar = s.calculatePayment(fields[2], hours)

			break
		}
	}

	if err := s.repo.Delete(req.PlatNomor); err != nil {
		return model.ExitVehicleResponse{}, err
	}

	return response, nil
}

func (s *serviceImpl) validateRegistration(plateNumber string, lines []string) error {
	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if fields[0] == plateNumber {
			return lib.NewErrDuplicate(fmt.Sprintf("plat_nomor %s already registered", plateNumber))
		}
	}
	return nil
}

func (s *serviceImpl) validateExit(plateNumber string, lines []string) error {
	found := false
	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		if fields[0] == plateNumber {
			found = true
			break
		}
	}

	if !found {
		return lib.NewErrNotFound(fmt.Sprintf("plat_nomor %s not found", plateNumber))
	}
	return nil
}

func (s *serviceImpl) generateParkingLot(lines []string) string {
	for i := 1; i <= lib.Cfg.MaxParkingLot; i++ {
		found := false
		for _, line := range lines {
			if line == "" {
				continue
			}

			fields := strings.Split(line, ",")
			if fields[3] == fmt.Sprintf("A%d", i) {
				found = true
				break
			}
		}

		if !found {
			return fmt.Sprintf("A%d", i)
		}
	}

	return ""
}

func (s *serviceImpl) calculateHoursParked(startTime, endTime string) int {
	t1, _ := time.Parse("2006-01-02 15:04", startTime)
	t2, _ := time.Parse("2006-01-02 15:04", endTime)
	duration := t2.Sub(t1)
	hours := int(duration.Hours()) + 1 // to handle partial hours as full
	return hours
}

func (s *serviceImpl) calculatePayment(vehicleType string, hours int) int {
	rate := 0

	switch vehicleType {
	case "SUV":
		rate = lib.Cfg.RateCarType.SUV
	case "MPV":
		rate = lib.Cfg.RateCarType.MPV
	}

	additionalRate := int(float64(rate) * 0.2)
	return rate + (hours-1)*additionalRate
}
