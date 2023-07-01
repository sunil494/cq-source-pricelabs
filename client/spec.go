package client

type Listing struct {
	ID      string `json:"id"`
	PMS     string `json:"pms"`
	API_KEY string `json:"api_key"`
}

type Spec struct {
	Listings []Listing `json:"listings"`
}
