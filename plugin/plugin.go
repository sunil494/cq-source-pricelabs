package plugin

import (
	"github.com/sunil494/cq-source-pricelabs/client"
	"github.com/sunil494/cq-source-pricelabs/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"PriceLabs",
		Version,
		schema.Tables{
			resources.ListingsTable(),
		},
		client.New,
	)
}
