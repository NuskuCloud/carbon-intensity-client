package carbon_intensity

// Define structures for Carbon Intensity API

type IntensityData struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Intensity struct {
		Forecast int    `json:"forecast"`
		Actual   int    `json:"actual"`
		Index    string `json:"index"`
	} `json:"intensity"`
}

type IntensityResponse struct {
	Data []IntensityData `json:"data"`
}

type DeterministicForecastData struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Intensity struct {
		Forecast int    `json:"forecast"`
		Index    string `json:"index"`
	} `json:"intensity"`
}

type DeterministicForecastResponse struct {
	Data []DeterministicForecastData `json:"data"`
}

type IntensityFactors struct {
	Biomass       int `json:"biomass"`
	Coal          int `json:"coal"`
	DutchImports  int `json:"dutchImports"`
	FrenchImports int `json:"frenchImports"`
	// Add other fields as necessary
}

type IntensityFactorsResponse struct {
	Data []IntensityFactors `json:"data"`
}

// CarbonIntensityResponse represents the response structure for the postcode endpoint
type CarbonIntensityResponse struct {
	Data []struct {
		From      string `json:"from"`
		To        string `json:"to"`
		Intensity struct {
			Forecast int    `json:"forecast"`
			Index    string `json:"index"`
		} `json:"intensity"`
	} `json:"data"`
}
