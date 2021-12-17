package model 

type Storage struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	PriceUSD     string `json:"price_usd"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

type StorageAdmin struct {
	AreaProvinsi string           `json:"area_provinsi"`
	Data         StorageAdminData `json:"data"`
}

type StorageAdminData struct {
	Min    int     `json:"min"`
	Max    int     `json:"max"`
	Median int     `json:"median"`
	Avg    float64 `json:"avg"`
}

type GetStoragesRequest struct {
	Rates float64
	Role  string
}
