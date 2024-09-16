package data

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
