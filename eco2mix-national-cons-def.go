package ego2mix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type ResponseConsDef struct {
	NHits      int               `json:"nhits"`
	Parameters ParametersConsDef `json:"parameters"`
	Records    []RecordConsDef   `json:"records"`
}

type ParametersConsDef struct {
	Dataset  string   `json:"dataset"`
	Q        string   `json:"q"`
	Rows     int      `json:"rows"`
	Start    int      `json:"start"`
	Sort     []string `json:"sort"`
	Facet    []string `json:"facet"`
	Format   string   `json:"format"`
	Timezone string   `json:"timezone"`
}

type RecordConsDef struct {
	DatasetID       string        `json:"datasetid"`
	RecordID        string        `json:"recordid"`
	Fields          FieldsConsDef `json:"fields"`
	RecordTimeStamp string        `json:"record_timestamp"`
}

type FieldsConsDef struct {
	Fioul                    int    `json:"fioul"`
	Hydraulique              int    `json:"hydraulique"`
	FioulAutres              int    `json:"fioul_autres"`
	PrevisionJ               int    `json:"prevision_j"`
	HydrauliqueStepTurbinage int    `json:"hydraulique_step_turbinage"`
	Nature                   string `json:"nature"`
	EchCommItalie            int    `json:"ech_comm_italie"`
	FioulCogen               int    `json:"fioul_cogen"`
	EchPhysiques             int    `json:"ech_physiques"`
	EchCommAngleterre        int    `json:"ech_comm_angleterre"`
	Charbon                  int    `json:"charbon"`
	Nucleaire                int    `json:"nucleaire"`
	DateHeure                string `json:"date_heure"`
	FioulTac                 int    `json:"fioul_tac"`
	Heure                    string `json:"heure"`
	Solaire                  int    `json:"solaire"`
	Perimetre                string `json:"perimetre"`
	EchCommEspagne           int    `json:"ech_comm_espagne"`
	BioenergiesBiogaz        int    `json:"bioenergies_biogaz"`
	Consommation             int    `json:"consommation"`
	Pompage                  int    `json:"pompage"`
	EchCommSuisse            int    `json:"ech_comm_suisse"`
	Date                     string `json:"date"`
	Gaz                      int    `json:"gaz"`
	Bioenergies              int    `json:"bioenergies"`
	GazCogen                 string `json:"gaz_cogen"`
	BioenergiesDechets       int    `json:"bioenergies_dechets"`
	TauxCo2                  int    `json:"taux_co2"`
	BioenergiesBiomasse      int    `json:"bioenergies_biomasse"`
	HydrauliqueFilEauEclusee int    `json:"hydraulique_fil_eau_eclusee"`
	GazCCG                   int    `json:"gaz_ccg"`
	EchCommAllemagneBelgique string `json:"ech_comm_allemagne_belgique"`
	HydrauliqueLacs          int    `json:"hydraulique_lacs"`
	GazTac                   int    `json:"gaz_tac"`
	GazAutres                int    `json:"gaz_autres"`
	Eolien                   int    `json:"eolien"`
	PrevisionJ1              int    `json:"prevision_j1"`
}

func (client *Eco2mixClient) FetchNationalFinalData(from time.Time, to time.Time, maxResults int) ([]FieldsConsDef, error) {
	params := url.Values{}
	params.Add("dataset", "eco2mix-national-cons-def")
	params.Add("facet", "nature")
	params.Add("facet", "date_heure")
	params.Add("start", "0")
	params.Add("rows", fmt.Sprintf("%d", maxResults))
	params.Add("sort", "date_heure")
	params.Add("q", fmt.Sprintf("date_heure:[%s TO %s] AND NOT #null(taux_co2)", from.Format("2006-01-02"), to.Format("2006-01-02")))
	queryString := params.Encode()

	resp, err := client.httpClient.Get(fmt.Sprintf(OPENDATASOFT_API_PATH, client.BaseUrl, queryString))
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %s", err)
	}
	// length := resp.ContentLength
	// fmt.Printf("resp content length: %d\n", length)
	// fmt.Printf("resp: %v\n", resp)

	// print curl equivalent command
	// fmt.Printf("curl -X GET %s\n", fmt.Sprintf(OPENDATASOFT_API_PATH, client.BaseUrl, queryString))
	// fmt.Printf("query string: %s\n", queryString)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: status=%s : body=%v", resp.Status, string(body))
	}
	// fmt.Printf("body: %s\n", body)
	var data ResponseConsDef
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %s", err)
	}

	var fields []FieldsConsDef
	for _, r := range data.Records {
		fields = append(fields, r.Fields)
	}

	return fields, nil
}
