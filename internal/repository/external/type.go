package external

const (
	storageURL      = "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"
	converterAPIKey = "547683d6c762119fc335"
	converterURL    = "https://free.currconv.com/api/v7/convert?q=%s&compact=ultra&apiKey=%s"
)

type Rates struct {
	IDRUSD float64 `json:"IDR_USD,omitempty"`
	USDIDR float64 `json:"USD_IDR,omitempty"`
}
