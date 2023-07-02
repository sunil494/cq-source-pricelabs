package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
	"github.com/sunil494/cq-source-pricelabs/internal/pricelabs"
)

type Client struct {
	Logger    zerolog.Logger
	PriceLabs *pricelabs.Client
	Spec      *Spec

	Listing PriceLabsConfigBlock
}

func (c *Client) ID() string {
	return fmt.Sprintf("pricelabs:%s", c.Listing.NAME)
}

func (c *Client) WithListing(listing PriceLabsConfigBlock) *Client {
	newC := *c
	newC.Logger = c.Logger.With().Str("listing", listing.NAME).Logger()
	newC.Listing = listing
	return &newC
}

func New(_ context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	c, err := pricelabs.NewClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		Logger:    logger,
		PriceLabs: c,
		Spec:      &pluginSpec,
	}, nil
}
