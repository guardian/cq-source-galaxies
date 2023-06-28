package client

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/guardian/cq-source-galaxies/store"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger zerolog.Logger
	Store  store.S3
}

func (c *Client) ID() string {
	return "guardian/galaxies"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("eu-west-1"),
		config.WithSharedConfigProfile("deployTools"),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config, %w", err)
	}

	client := s3.NewFromConfig(cfg)
	store := store.New(client, pluginSpec.GalaxiesBucketName, "galaxies.gutools.co.uk/data")

	return &Client{
		Logger: logger,
		Store:  store,
	}, nil
}
