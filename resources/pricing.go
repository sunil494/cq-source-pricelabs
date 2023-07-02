package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/sunil494/cq-source-pricelabs/client"
	"github.com/sunil494/cq-source-pricelabs/internal/pricelabs"
)

func PricingTable() *schema.Table {
	return &schema.Table{
		Name:      "pricelabs_listing_pricing",
		Resolver:  fetchPriceLabsData,
		Multiplex: client.ListingMultiplex,
		Transform: transformers.TransformWithStruct(&pricelabs.PriceLabsPricing{}),
	}
}

func fetchPriceLabsData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	listing := c.Listing
	pricing, err := c.PriceLabs.GetPriceLabs(listing.API_KEY, listing.ID, listing.PMS)
	if err != nil {
		return err
	}
	for _, price := range pricing.Data {
		res <- price
	}
	return nil
}
