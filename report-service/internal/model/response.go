package model

type CountByTypeResponse struct {
	JumlahKendaraan int `json:"jumlah_kendaraan"`
}

type ListByColorResponse struct {
	PlatNomor []string `json:"plat_nomor"`
}
