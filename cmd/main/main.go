package main

import (
	"fmt"
	"github.com/nuskucloud/carbon-intensity-client/pkg/carbon_intensity"
	"log"
	"time"
)

func main() {
	client := carbon_intensity.NewClient()

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
	//forecast, err := client.GetDeterministicForecast(oneMonthAgo, carbon_intensity.TwentyFourHoursAfter)
	//forecast, err := client.GetDeterministicForecast(oneMonthAgo, carbon_intensity.FortyEightHoursAfter)
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

	postcodeIntensity, err := client.GetCarbonIntensityByPostcode("DE24") // Do not use the inward code (ie, the last 3 characters)
	if err != nil {
		log.Fatalf("Error fetching intensity factors: %v", err)
	}
	fmt.Printf("Intensity Factors Data: %+v\n", postcodeIntensity)

	regionalIntensity, err := client.GetRegionalIntensityFactors(carbon_intensity.England)
	if err != nil {
		log.Fatalf("Error fetching intensity factors: %v", err)
	}
	fmt.Printf("Intensity Factors Data: %+v\n", regionalIntensity)

	generationMix, err := client.GetGenerationMix()
	if err != nil {
		log.Fatalf("Error fetching intensity factors: %v", err)
	}
	fmt.Printf("Intensity Factors Data: %+v\n", generationMix)
}
