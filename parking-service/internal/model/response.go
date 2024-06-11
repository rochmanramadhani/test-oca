package model

type RegisterVehicleResponse struct {
	PlatNomor    string `json:"plat_nomor"`
	ParkingLot   string `json:"parking_lot"`
	TanggalMasuk string `json:"tanggal_masuk"`
}

type ExitVehicleResponse struct {
	PlatNomor     string `json:"plat_nomor"`
	TanggalMasuk  string `json:"tanggal_masuk"`
	TanggalKeluar string `json:"tanggal_keluar"`
	JumlahBayar   int    `json:"jumlah_bayar"`
}
