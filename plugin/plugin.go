package plugin

import (
	"github.com/sunil/cq-source-test/client"
	"github.com/sunil/cq-source-test/resources"

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
			resources.SampleTable(),
		},
		client.New,
	)
}
