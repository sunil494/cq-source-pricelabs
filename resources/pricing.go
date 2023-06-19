package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/sunil494/cq-source-test/client"
	"github.com/sunil494/cq-source-test/internal/pricelabs"
	"golang.org/x/sync/errgroup"
)

func PricingTable() *schema.Table {
	return &schema.Table{
		Name:      "pricelabs_listing_pricing",
		Resolver:  fetchPriceLabsData,
		Transform: transformers.TransformWithStruct(&pricelabs.PriceLabsPricing{}),
	}
}

func fetchPriceLabsData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	pricing, err := c.PriceLabs.GetPriceLabs(c.Spec.ApiKey, "1", "2")
	if err != nil {
		return err
	}
	for _, price := range pricing.Data {
		res <- price
	}
	g := errgroup.Group{}
	g.SetLimit(10)
	return g.Wait()
}
