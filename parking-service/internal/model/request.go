package model

import (
	"errors"
	"lib"
)

type RegisterVehicleRequest struct {
	PlatNomor string      `json:"plat_nomor"`
	Warna     lib.Color   `json:"warna"`
	Tipe      lib.CarType `json:"tipe"`
}

func (r RegisterVehicleRequest) Validate() error {
	if r.PlatNomor == "" {
		return errors.New("plat_nomor cannot be empty")
	} else {
		if !lib.CheckLicensePlateFormat(r.PlatNomor) {
			return errors.New("plat_nomor format is invalid")
		}
	}

	if r.Warna == "" {
		return errors.New("warna cannot be empty")
	} else {
		_, err := lib.Color.Parse(r.Warna, string(r.Warna))
		if err != nil {
			return err
		}
	}

	if r.Tipe == "" {
		return errors.New("tipe cannot be empty")
	} else {
		_, err := lib.CarType.Parse(r.Tipe, string(r.Tipe))
		if err != nil {
			return err
		}
	}

	return nil
}

type ExitVehicleRequest struct {
	PlatNomor string `json:"plat_nomor"`
}

func (r ExitVehicleRequest) Validate() error {
	if r.PlatNomor == "" {
		return errors.New("plat_nomor cannot be empty")
	} else {
		if !lib.CheckLicensePlateFormat(r.PlatNomor) {
			return errors.New("plat_nomor format is invalid")
		}
	}

	return nil
}
