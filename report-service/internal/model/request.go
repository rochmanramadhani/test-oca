package model

import "lib"

type CountByTypeRequest struct {
	Tipe lib.CarType
}

func (r CountByTypeRequest) Validate() error {
	if r.Tipe == "" {
		return lib.NewErrBadRequest("q is required")
	} else {
		_, err := lib.CarType.Parse(r.Tipe, string(r.Tipe))
		if err != nil {
			return err
		}
	}
	return nil
}

type ListByColorRequest struct {
	Warna lib.Color
}

func (r ListByColorRequest) Validate() error {
	if r.Warna == "" {
		return lib.NewErrBadRequest("q is required")
	} else {
		_, err := lib.Color.Parse(r.Warna, string(r.Warna))
		if err != nil {
			return err
		}
	}
	return nil
}
