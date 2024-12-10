package data

import (
	"context"
	"errors"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type School struct {
	Udise           string `json:"udise"`
	School_name     string `json:"school_name"`
	School_location string `json:"school_location"`
	School_category string `json:"school_category"`
	School_type     string `json:"school_type"`
	Village         string `json:"village"`
	Block           string `json:"block"`
	Cluster         string `json:"cluster"`
	District        string `json:"district"`
	State           string `json:"state"`
	Lat             string `json:"lat"`
	Long            string `json:"long"`
	Donor           string `json:"donor"`
}

var (
	ErrRecordNotFound = errors.New("Record not found.")
	ErrQueryFailed    = errors.New("Failed to run query.")
	ErrInvalidUDISE   = errors.New("Invalid UDISE.")
	ErrIteratorFailed = errors.New("Row iterator failed.")
)

func isValidUdise(udise string) bool {
	// Check if the UDISE code is 11 digits long
	if len(udise) != 11 {
		return false
	}

	// If all checks pass, UDISE is valid
	return true
}

func Get(id string, bqClient *bigquery.Client, ctx context.Context) (*School, error) {
	udise := id
	if !isValidUdise(udise) {
		return nil, ErrInvalidUDISE
	}

	q := bqClient.Query(`
	SELECT *
	FROM afe-bot.quest_schools_matrix.udise_superset_prod
	WHERE udise = @udise
	`)

	q.Parameters = []bigquery.QueryParameter{
		{Name: "udise", Value: udise},
	}

	it, err := q.Read(ctx)
	if err != nil {
		return nil, ErrQueryFailed
	}

	for {
		var row map[string]bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			// Query returned no rows
			return nil, ErrRecordNotFound
		}
		if err != nil {
			return nil, ErrIteratorFailed
		}

		school := &School{}

		if val, ok := row["udise"]; ok {
			school.Udise, _ = val.(string)
		}

		if val, ok := row["school_name"]; ok {
			school.School_name, _ = val.(string)
		}

		if val, ok := row["school_location"]; ok {
			school.School_location, _ = val.(string)
		}

		if val, ok := row["school_category"]; ok {
			school.School_category, _ = val.(string)
		}

		if val, ok := row["school_type"]; ok {
			school.School_type, _ = val.(string)
		}

		if val, ok := row["village"]; ok {
			school.Village, _ = val.(string)
		}

		if val, ok := row["block"]; ok {
			school.Block, _ = val.(string)
		}

		if val, ok := row["cluster"]; ok {
			school.Cluster, _ = val.(string)
		}

		if val, ok := row["district"]; ok {
			school.District, _ = val.(string)
		}

		if val, ok := row["state"]; ok {
			school.State, _ = val.(string)
		}

		if val, ok := row["lat"]; ok {
			school.Lat, _ = val.(string)
		}

		if val, ok := row["long"]; ok {
			school.Long, _ = val.(string)
		}

		if val, ok := row["donor"]; ok {
			school.Donor, _ = val.(string)
		}

		return school, nil
	}
}
