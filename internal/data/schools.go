package data

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type School struct {
	Udise           string `json:"udise"`
	School_name     string `json:"school_name"`
	School_area     string `json:"school_area"`
	Village_or_town string `json:"village_or_town"`
	Cluster         string `json:"cluster"`
	Block           string `json:"block"`
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
)

func isValidUdise(udise string) bool {

	// Check if the UDISE code is 11 digits long
	if len(udise) != 11 {
		return false
	}

	// First two digits: state (must be between 01 and 37 for valid Indian states)
	state, err := strconv.Atoi(udise[:2])
	if err != nil || state < 1 || state > 37 {
		return false
	}

	// Next three digits: district (must be between 001 and 999)
	district, err := strconv.Atoi(udise[2:5])
	if err != nil || district < 1 || district > 999 {
		return false
	}

	// Next three digits: block (must be between 001 and 999)
	block, err := strconv.Atoi(udise[5:8])
	if err != nil || block < 1 || block > 999 {
		return false
	}

	// Last digit: check digit (should be a numeric digit)
	checkDigit := udise[10]
	if !unicode.IsDigit(rune(checkDigit)) {
		return false
	}

	// If all checks pass, UDISE is valid
	return true
}

func Get(id int64, bqClient *bigquery.Client, ctx context.Context) (*School, error) {

	udise := strconv.FormatInt(id, 10)

	if !isValidUdise(udise) {
		return nil, ErrRecordNotFound
	}

	q := bqClient.Query(`
	SELECT *
	FROM afe-bot.quest_schools_matrix.school_details_fact
	WHERE udise = @udise
	`)

	q.Parameters = []bigquery.QueryParameter{
		{Name: "udise", Value: udise},
	}

	it, err := q.Read(ctx)
	if err != nil {
		return nil, ErrQueryFailed

	}

	/**
	for {
		var row map[string]bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			// No more rows
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate through query results: %w", err)
		}

		// Return the first found record
		return row, nil
	}

	**/

	for {
		var row map[string]bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			// No more rows
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate through query results: %w", err)
		}

		// Map BigQuery row to School struct
		school := &School{}

		if val, ok := row["udise"]; ok {
			school.Udise, _ = val.(string)
		}

		if val, ok := row["school_name"]; ok {
			school.School_name, _ = val.(string)
		}

		if val, ok := row["school_area"]; ok {
			school.School_area, _ = val.(string)
		}

		if val, ok := row["village_or_town"]; ok {
			school.Village_or_town, _ = val.(string)
		}

		if val, ok := row["cluster"]; ok {
			school.Cluster, _ = val.(string)
		}

		if val, ok := row["block"]; ok {
			school.Block, _ = val.(string)
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

		// Return the first found record
		fmt.Printf("Record: %+v\n", school)
		return school, nil
	}

}
