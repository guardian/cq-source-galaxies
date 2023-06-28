package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/guardian/cq-source-galaxies/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
