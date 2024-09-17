package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type School struct {
	SchoolName                 string `json:"school_name"`
	VillageOrTown              string `json:"village_or_town"`
	Cluster                    string `json:"cluster"`
	Block                      string `json:"block"`
	District                   string `json:"district"`
	State                      string `json:"state"`
	UDISECode                  int64  `json:"udise_code"`
	Building                   string `json:"building"`
	ClassRooms                 int    `json:"class_rooms"`
	BoysToilet                 int    `json:"boys_toilet"`
	GirlsToilet                int    `json:"girls_toilet"`
	ComputerAidedLearning      string `json:"computer_aided_learning"`
	Electricity                string `json:"electricity"`
	Wall                       string `json:"wall"`
	Library                    string `json:"library"`
	Playground                 string `json:"playground"`
	BooksInLibrary             int    `json:"books_in_library"`
	DrinkingWater              string `json:"drinking_water"`
	RampsForDisabled           string `json:"ramps_for_disabled"`
	Computers                  int    `json:"computers"`
	InstructionMedium          string `json:"instruction_medium"`
	MaleTeachers               int    `json:"male_teachers"`
	PrePrimarySectionAvailable string `json:"pre_primary_section_available"`
	BoardForClass10th          string `json:"board_for_class_10th"`
	SchoolType                 string `json:"school_type"`
	Classes                    string `json:"classes"`
	FemaleTeachers             int    `json:"female_teachers"`
	PrePrimaryTeachers         int    `json:"pre_primary_teachers"`
	BoardForClass10Plus2       string `json:"board_for_class_10_plus_2"`
	Meal                       string `json:"meal"`
	Establishment              string `json:"establishment"`
	SchoolArea                 string `json:"school_area"`
	SchoolShiftedToNewPlace    string `json:"school_shifted_to_new_place"`
	HeadTeachers               int    `json:"head_teachers"`
	HeadTeacherName            string `json:"head_teacher_name"`
	IsSchoolResidential        string `json:"is_school_residential"`
	ResidentialType            string `json:"residential_type"`
	TotalTeachers              int    `json:"total_teachers"`
	ContractTeachers           int    `json:"contract_teachers"`
	Management                 string `json:"management"`
	Latitude                   string `json:"latitude"`
	Longitude                  string `json:"longitude"`
}

var (
	ErrRecordNotFound = errors.New("record not found")
)

func Get(id int64, bqClient *bigquery.Client, ctx context.Context) (map[string]bigquery.Value, error) {

	// [TODO] : Decide if we want to fetch project, dataset or table from command line args
	q := bqClient.Query(`
	SELECT *
	FROM afe-bot.quest_schools_matrix.school_details_fact
	WHERE ` + "`UDISE Code`" + ` = @id
	`)

	q.Parameters = []bigquery.QueryParameter{
		{Name: "id", Value: strconv.FormatInt(id, 10)},
	}

	it, err := q.Read(ctx)
	if err != nil {
		log.Fatalf("Failed to run query: %v", err)
	}

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

	/** [TODO] : Decide if dat should be returned as School struct

	for {
		var row map[string]bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			// No more rows
			return nil, ErrRecordNotFound
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate through query results: %w", err)
		}

		// Map BigQuery row to School struct
		school := &School{}

		if val, ok := row["School Name"]; ok {
			school.SchoolName, _ = val.(string)
		}
		if val, ok := row["Village or Town"]; ok {
			school.VillageOrTown, _ = val.(string)
		}
		if val, ok := row["Cluster"]; ok {
			school.Cluster, _ = val.(string)
		}
		if val, ok := row["Block"]; ok {
			school.Block, _ = val.(string)
		}
		if val, ok := row["District"]; ok {
			school.District, _ = val.(string)
		}
		if val, ok := row["State"]; ok {
			school.State, _ = val.(string)
		}
		if val, ok := row["UDISE Code"]; ok {
			if udiseCode, err := strconv.ParseInt(val.(string), 10, 64); err == nil {
				school.UDISECode = udiseCode
			}
		}

		// Return the first found record
		fmt.Printf("Record: %+v\n", school)
		return school, nil
	}

	**/
}
