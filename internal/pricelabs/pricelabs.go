package pricelabs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode/utf8"
)

const defaultURL = "https://api.pricelabs.co/v1/listing_prices"

type PriceLabs struct {
	Id                string             `json:"Id"`
	Pms               string             `json:"pms"`
	Currency          string             `json:"currency"`
	Last_refreshed_at string             `json:"last_refreshed_at"`
	Data              []PriceLabsPricing `json:"data"`
}

type PriceLabsListing struct {
	Id                string `json:"Id"`
	Pms               string `json:"pms"`
	Currency          string `json:"currency"`
	Last_refreshed_at string `json:"last_refreshed_at"`
}

type PriceLabsPricing struct {
	Date  string `json:"date"`
	Price int    `json:"price"`
}

type Client struct {
	baseURL string
	client  *http.Client
}

type Option func(*Client)

func WithBaseURL(uri string) Option {
	return func(c *Client) {
		c.baseURL = uri
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		baseURL: defaultURL,
		client:  http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[1 : len(s)-size]
}

func (c *Client) GetPriceLabs(api_key string, listing_id string, pms string) (*PriceLabs, error) {

	url := "https://api.pricelabs.co/v1/listing_prices"
	method := "POST"

	payload := strings.NewReader(`{
    "listings": [
        {
        "id": "` + listing_id + `",
        "pms": "` + pms + `"
        }
    ]
    }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("x-api-key", api_key)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	foo2 := PriceLabs{}

	body, err := io.ReadAll(res.Body)
	if err := json.Unmarshal([]byte(trimLastChar(string(body))), &foo2); err != nil {
		fmt.Println(err)
		return &foo2, err
	}
	return &foo2, err
}
