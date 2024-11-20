package carbon_intensity

type GenerationMixResponse struct {
	Data GenerationMixData `json:"data"`
}

type GenerationMixData struct {
	From          string          `json:"from"`
	To            string          `json:"to"`
	GenerationMix []GenerationMix `json:"generationmix"`
}

type GenerationMix struct {
	Fuel    string  `json:"fuel"`
	PerCent float64 `json:"perc"`
}

func (c *CarbonIntensityClient) GetGenerationMix() (*GenerationMixResponse, error) {
	var generationMix GenerationMixResponse
	err := c.MakeRequest("/generation", nil, &generationMix)
	return &generationMix, err
}
