package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/guardian/cq-source-galaxies/client"
	"github.com/guardian/cq-source-galaxies/resources"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"guardian-galaxies",
		Version,
		schema.Tables{
			resources.PeopleTable(),
			resources.StreamsTable(),
			resources.TeamsTable(),
		},
		client.New,
	)
}
