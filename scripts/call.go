package main

import (
	"fmt"
	"time"

	ego2mix "github.com/timebis/ego2mix"
)

func main() {
	client := ego2mix.NewEco2mixClient("", nil)

	// round to day timenow
	from := time.Now().Add(-24 * time.Hour).Round(24 * time.Hour)
	to := from.Add(24 * time.Hour)

	fmt.Printf("from: %s\n", from)
	fmt.Printf("to: %s\n", to)

	records, err := client.FetchNationalRealTimeData(from, to, 100)
	if err != nil {
		panic(err)
	}

	fmt.Printf("records: %v\n\n", records)
	fmt.Printf("Intensité carbone à %s le %s en France: %d gCO2eq / kWh\n", records[0].Heure, records[0].Date, records[0].TauxCo2)

	// find Co2 intensity at 6:38
	today6h38 := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 6, 38, 0, 0, time.Local)

	co2Intensity6h38, err := ego2mix.FindClosestRecord(records, today6h38, 20*time.Minute)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Intensité carbone à 6h38 le %s en France: %d gCO2eq / kWh\n", co2Intensity6h38.Date, co2Intensity6h38.TauxCo2)

}
