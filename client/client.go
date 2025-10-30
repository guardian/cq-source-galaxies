package client

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/guardian/cq-source-galaxies/store"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	Spec   Spec
	Store  store.S3
}

func (c *Client) ID() string {
	return "guardian/galaxies"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	// Loads credentials from the default credential chain.
	// Locally, set the AWS_PROFILE environment variable, or run `make serve`.
	// See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials.
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("eu-west-1"),
	)

	if err != nil {
		log.Fatalf("unable to load AWS config, %v", err)
	}

	client := s3.NewFromConfig(cfg)
	s3Store := store.New(client, s.GalaxiesBucketName, "galaxies.gutools.co.uk/data")

	return Client{
		logger: logger,
		Spec:   *s,
		Store:  s3Store,
	}, nil
}
