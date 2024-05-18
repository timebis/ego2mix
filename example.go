package ego2mix

import (
	"fmt"
	"log"
	"time"
)

func main() {
	client := NewEco2mixClient("", nil)
	from := time.Now().Add(-72 * time.Hour)
	to := from.Add(24 * time.Hour)
	records, err := client.FetchNationalRealTimeData(from, to, 10) // we want only the last result
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Intensité carbone à %s en France: %d gCO2eq / kWh\n", records[0].DateHeure, records[0].TauxCo2)
}
