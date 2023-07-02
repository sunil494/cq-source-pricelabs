package client

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func ListingMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for index := range client.Spec.Listings {
		l = append(l, client.WithListing(client.Spec.Listings[index]))
	}
	return l
}
