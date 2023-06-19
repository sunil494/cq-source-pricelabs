package plugin

import (
	"github.com/sunil494/cq-source-test/client"
	"github.com/sunil494/cq-source-test/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"sunil-test",
		Version,
		schema.Tables{
			resources.ListingsTable(),
		},
		client.New,
	)
}
