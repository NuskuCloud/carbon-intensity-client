package carbon_intensity

import "fmt"

const (
	TwentyFourHoursBefore = "pt24h"
	TwentyFourHoursAfter  = "fw24h"
	FortyEightHoursAfter  = "fw48h"

	England  = "England"
	Scotland = "Scotland"
	Wales    = "Wales"
)

func (c *CarbonIntensityClient) GetIntensity() (*IntensityResponse, error) {
	var intensity IntensityResponse
	err := c.MakeRequest("/intensity", nil, &intensity)
	return &intensity, err
}

func (c *CarbonIntensityClient) GetDeterministicForecast(from string, to string) (*DeterministicForecastResponse, error) {
	var forecast DeterministicForecastResponse

	err := c.MakeRequest(fmt.Sprintf("/intensity/%s/%s", from, to), nil, &forecast)
	return &forecast, err
}

func (c *CarbonIntensityClient) GetIntensityFactors() (*IntensityFactorsResponse, error) {
	var factors IntensityFactorsResponse
	err := c.MakeRequest("/intensity/factors", nil, &factors)
	return &factors, err
}

func (c *CarbonIntensityClient) GetRegionalIntensityFactors(region string) (*IntensityFactorsResponse, error) {
	var factors IntensityFactorsResponse
	err := c.MakeRequest("/regional/"+region, nil, &factors)
	return &factors, err
}
