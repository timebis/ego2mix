package ego2mix

import (
	"testing"
	"time"
)

func TestFetchNationalRealTimeData(t *testing.T) {
	from := time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 7, 5, 0, 0, 0, 0, time.UTC)

	c := NewEco2mixClient("", nil)

	res, err := c.FetchNationalRealTimeData(from, to, 3)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	if len(res) != 3 {
		t.Errorf("expected 3 records, got %d", len(res))
	}

	expectedTauxCo2 := []int64{32, 32, 31}
	for i := range expectedTauxCo2 {
		if res[i].TauxCo2 != expectedTauxCo2[i] {
			t.Errorf("record %d: expected value %d, got %d", i, expectedTauxCo2[i], res[i].TauxCo2)

		}
	}
}
