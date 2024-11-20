# Go client for the carbon intensity data published by National Energy System Operator (NESO)

# Carbon Intensity API Client

A Go client for interacting with the Carbon Intensity API. This client provides functions to fetch various types of carbon intensity data, including generation mix, deterministic forecasts, intensity factors, and postcode-based intensity.

The API reference is available here - https://carbon-intensity.github.io/api-definitions/

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

```sh
go get -u github.com/nuskucloud/carbon-intensity-api/pkg/carbon_intensity
```

### Usage

You can do a `git clone` run the provided `main.go` file to see the client in action. It fetches various data from the Carbon Intensity API and prints it to the console.

To run the main program:

```sh
go run main.go
```

## Project Structure

- `main.go`: Entry point of the application. Demonstrates how to use the `carbon_intensity` package to fetch data from the API.
- `pkg/carbon_intensity/api_client.go`: Implements the HTTP client for making requests to the Carbon Intensity API.
- `pkg/carbon_intensity/service_intensity.go`: Provides functions to fetch general carbon intensity data.
- `pkg/carbon_intensity/service_postcode.go`: Provides functions to fetch carbon intensity data based on postcodes.
- `pkg/carbon_intensity/service_generation.go`: Provides functions to fetch the generation mix data.
- `pkg/carbon_intensity/models.go`: Contains the data models used throughout the client.

## Example

```go
package main

import (
	carbonIntensity "github.com/nuskucloud/carbon-intensity-api/pkg/carbon_intensity"
    "fmt"
    "log"
    "time"
)

func main() {
    client := carbonIntensity.NewClient()
    
    // Fetch Intensity
    intensity, err := client.GetIntensity()
    if err != nil {
        log.Fatalf("Error fetching intensity: %v", err)
    }
    fmt.Printf("Intensity Data: %+v\n", intensity)
    
    // Fetch Historical Deterministic Forecast
    oneMonthAgo := time.Now().AddDate(0, -1, 0).Format(time.RFC3339)
    threeDaysAfterOneMonthAgo := time.Now().AddDate(0, 0, -2).Format(time.RFC3339)
    forecast, err := client.GetDeterministicForecast(oneMonthAgo, threeDaysAfterOneMonthAgo)
    if err != nil {
        log.Fatalf("Error fetching deterministic forecast: %v", err)
    }
    fmt.Printf("Deterministic Forecast Data: %+v\n", forecast)
    
    // Fetch standard chunk Deterministic Forecast
    forecast, err = client.GetDeterministicForecast(oneMonthAgo, carbon_intensity.TwentyFourHoursBefore)
    if err != nil {
        log.Fatalf("Error fetching deterministic forecast: %v", err)
    }
    fmt.Printf("Deterministic Forecast Data: %+v\n", forecast)
    
    // Fetch Intensity Factors
    factors, err := client.GetIntensityFactors()
    if err != nil {
        log.Fatalf("Error fetching intensity factors: %v", err)
    }
    fmt.Printf("Intensity Factors Data: %+v\n", factors)

    // Fetch Carbon Intensity by Postcode
    postcodeIntensity, err := client.GetCarbonIntensityByPostcode("DE24")
    if err != nil {
        log.Fatalf("Error fetching intensity by postcode: %v", err)
    }
    fmt.Printf("Postcode Intensity Data: %+v\n", postcodeIntensity)
    
    // Fetch Regional Intensity Factors
    regionalIntensity, err := client.GetRegionalIntensityFactors(carbon_intensity.England)
    if err != nil {
        log.Fatalf("Error fetching regional intensity factors: %v", err)
    }
    fmt.Printf("Regional Intensity Data: %+v\n", regionalIntensity)
    
    // Fetch Generation Mix
    generationMix, err := client.GetGenerationMix()
    if err != nil {
        log.Fatalf("Error fetching generation mix: %v", err)
    }
    fmt.Printf("Generation Mix Data: %+v\n", generationMix)
}
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

For any questions, feel free to reach out.

## License

This project is licensed under the MIT License.