package carbon_intensity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.carbonintensity.org.uk"

type CarbonIntensityClient struct {
	httpClient *http.Client
}

// NewClient creates a new CarbonIntensityClient
func NewClient() *CarbonIntensityClient {
	return &CarbonIntensityClient{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// MakeRequest is a helper function to send API requests
func (c *CarbonIntensityClient) MakeRequest(endpoint string, params map[string]string, target interface{}) error {
	base, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	base.Path += endpoint

	// Add query parameters if any
	if params != nil {
		query := base.Query()
		for key, value := range params {
			query.Set(key, value)
		}
		base.RawQuery = query.Encode()
	}

	resp, err := c.httpClient.Get(base.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received status code %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
