package pets

type Pet struct {
	Species          string `json:"species,omitempty"`
	Breed            string `json:"breed,omitempty"`
	MinWeight        int    `json:"min_weight,omitempty"`
	MaxWeight        int    `json:"max_weight,omitempty"`
	AverageWeight    int    `json:"average_weight,omitempty"`
	Weight           int    `json:"weight,omitempty"`
	Description      string `json:"description,omitempty"`
	Lifespan         string `json:"lifespan,omitempty"`
	GeographicOrigin string `json:"geographic_origin,omitempty"`
	Color            string `json:"color,omitempty"`
	Age              int    `json:"age,omitempty"`
	AgeEstimated     bool   `json:"age_estimated,omitempty"`
}
