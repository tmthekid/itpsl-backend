package main

type Response struct {
	Message string `json:"message"`
}

func response(message string) Response {
	return Response{Message: message}
}

type ConditionPayload struct {
	ConditionId int     `json:"condition_id"`
	Condition   string  `json:"condition"`
	Amount      float64 `json:"amount"`
	Percentage  float64 `json:"percentage"`
}

type ExpenditurePayload struct {
	ExpenditureId int     `json:"expenditure_id"`
	Expenditure   string  `json:"expenditure"`
	Amount        float64 `json:"amount"`
}

type DataPayload struct {
	CityValue         int                  `json:"city_value"`
	TaxRate           int                  `json:"tax_rate"`
	BondsAmount       int                  `json:"bonds_amount"`
	BondsInterestRate int                  `json:"bonds_interest_rate"`
	Conditions        []ConditionPayload   `json:"conditions"`
	Expenditures      []ExpenditurePayload `json:"expenditures"`
	Year              int                  `json:"year"`
	Growth            float64              `json:"growth"`
	ScoreId           int                  `json:"score_id"`
	ScoreCategory     string               `json:"score_category"`
}

type Payload struct {
	Name        string        `json:"name"`
	Age         string        `json:"age"`
	Proffesion  string        `json:"proffesion"`
	Gender      string        `json:"gender"`
	FinalGrowth string        `json:"final_growth"`
	Data        []DataPayload `json:"data"`
}

type Result struct {
	Id                              int8    `json:"id"`
	Name                            string  `json:"name"`
	Proffesion                      string  `json:"proffesion"`
	Gender                          string  `json:"gender"`
	Age                             string  `json:"age"`
	Year                            int     `json:"year"`
	Growth                          float32 `json:"growth"`
	ScoreCategory                   string  `json:"score_category"`
	CityValue                       float32 `json:"city_value"`
	TaxRate                         float32 `json:"tax_rate"`
	BondsAmount                     float32 `json:"bonds_amount"`
	BondsInterestRate               float32 `json:"bonds_interest_rate"`
	Drought                         float32 `json:"drought"`
	Flood                           float32 `json:"flood"`
	Landslide                       float32 `json:"landslide"`
	Tsunami                         float32 `json:"tsumami"`
	SocialUnrest                    float32 `json:"socrial_unrest"`
	Normal                          float32 `json:"normal"`
	PhysicalDevelopment             float32 `json:"physical_development"`
	EcoSystemPreservationProtection float32 `json:"eco_system_preservation_&_protection"`
	EconomicDevelopment             float32 `json:"economic_development"`
	PublicUtilitiesTransport        float32 `json:"public_utilities_&_transport"`
	WastePollutionManagement        float32 `json:"waste_&_pollution_management"`
	EducationHealthcare             float32 `json:"education_&_healthcare"`
	CulturalHeritageManagement      float32 `json:"cultural_&_heritage_management"`
	SocialSupport                   float32 `json:"social_support"`
	DisasterMitigationManagement    float32 `json:"disaster_mitigation_&_management"`
	ResearchInnovationDevelopment   float32 `json:"research_innovation_&_development"`
	MonitoringManagement            float32 `json:"monitoring_&_management"`
	FinalGrowth                     string  `json:"final_growth"`
}