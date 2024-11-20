package carbon_intensity

import (
	"fmt"
)

// GetCarbonIntensityByPostcode fetches carbon intensity data for a given postcode
func (c *CarbonIntensityClient) GetCarbonIntensityByPostcode(postcode string) (*CarbonIntensityResponse, error) {
	var response CarbonIntensityResponse
	endpoint := fmt.Sprintf("/regional/postcode/%s", postcode)
	err := c.MakeRequest(endpoint, nil, &response)
	return &response, err
}
