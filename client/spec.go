package client

type Spec struct {
	Listings []PriceLabsConfigBlock `json:"listings"`
}

type PriceLabsConfigBlock struct {
	NAME    string `json:"name"`
	ID      string `json:"id"`
	PMS     string `json:"pms"`
	API_KEY string `json:"api_key"`
}
