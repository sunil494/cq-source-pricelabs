package resources

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/sunil494/cq-source-pricelabs/client"
	"github.com/sunil494/cq-source-pricelabs/internal/pricelabs"
)

func ListingsTable() *schema.Table {
	return &schema.Table{
		Name:      "pricelabs_listings",
		Resolver:  fetchPriceLabs,
		Transform: transformers.TransformWithStruct(&pricelabs.PriceLabsListing{}),
		Relations: []*schema.Table{
			PricingTable(),
		},
	}
}

func fetchPriceLabs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	for _, listing := range c.Spec.Listings {
		listings, err := c.PriceLabs.GetPriceLabs(listing.API_KEY, listing.ID, listing.PMS)
		fmt.Printf("%+v\n", listings)
		if err != nil {
			return err
		}
		res <- listings
	}
	return nil
}
