package resources

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/guardian/cq-source-galaxies/client"
)

// These models reflect the Galaxies data models and will need updating if the
// underlying models change.

type Team struct {
	ID                 string `json:"teamId"`
	Name               string `json:"teamName"`
	Description        string `json:"teamDescription"`
	ContactEmail       string `json:"teamContactEmail"`
	GoogleChatSpaceKey string `json:"teamGoogleChatSpaceKey"`
	PrimaryGithubTeam  string `json:"teamPrimaryGithubTeam"`
}

type Stream struct {
	ID          string   `json:"streamId"`
	Name        string   `json:"streamName"`
	Description string   `json:"streamDescription"`
	Members     []string `json:"streamMembers"`
}

type Person struct {
	Name    string   `json:"name"`
	EmailID string   `json:"emailId"`
	Role    string   `json:"role"`
	Teams   []string `json:"teams"`
	Streams []string `json:"streams"`
}

func TeamsTable() *schema.Table {
	return &schema.Table{
		Name:      "galaxies_teams_table",
		Resolver:  fetchTeams,
		Transform: transformers.TransformWithStruct(&Team{}),
	}
}

func PeopleTable() *schema.Table {
	return &schema.Table{
		Name:      "galaxies_people_table",
		Resolver:  fetchPeople,
		Transform: transformers.TransformWithStruct(&Person{}),
	}
}

func StreamsTable() *schema.Table {
	return &schema.Table{
		Name:      "galaxies_streams_table",
		Resolver:  fetchStreams,
		Transform: transformers.TransformWithStruct(&Stream{}),
	}
}

func fetchPeople(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	store := c.Store

	data, err := store.Get("people.json")
	if err != nil {
		return err
	}

	var records []Person
	err = json.Unmarshal(data, &records)
	if err != nil {
		return err
	}

	for _, record := range records {
		res <- record
	}

	return nil

}

func fetchTeams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	store := c.Store

	data, err := store.Get("teams.json")
	if err != nil {
		return err
	}

	var records map[string]Team
	err = json.Unmarshal(data, &records)
	if err != nil {
		return err
	}

	for ID, record := range records {
		record.ID = ID // Galaxies team data uses the map keys as the ID, so pull this out.
		res <- record
	}

	return nil
}

func fetchStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	store := c.Store

	data, err := store.Get("streams.json")
	if err != nil {
		return err
	}

	var records map[string]Stream
	err = json.Unmarshal(data, &records)
	if err != nil {
		return err
	}

	for ID, record := range records {
		record.ID = ID // Galaxies streams data uses the map keys as the ID, so pull this out.
		res <- record
	}

	return nil
}
