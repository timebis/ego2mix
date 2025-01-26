package ego2mix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	OPENDATASOFT_API_DEFINITIVE_PATH = `%s/api/explore/v2.1/catalog/datasets/eco2mix-national-cons-def/records?%s`
)

type NationalDefinitiveRecord struct {
	TotalCount int                        `json:"total_count"`
	Results    []NationalDefinitiveFields `json:"results"`
}
type NationalDefinitiveFields struct {
	Bioenergies              int64  `json:"bioenergies"`                 // Bioénergies (MW)
	BioenergiesBiogaz        int64  `json:"bioenergies_biogaz"`          // Bioénergies - Biogaz (MW) - Production bioénergies réalisée à partir du biogaz
	BioenergiesBiomasse      int64  `json:"bioenergies_biomasse"`        // Bioénergies - Biomasse (MW) - Production bioénergies réalisée à partir de la biomasse
	BioenergiesDechets       int64  `json:"bioenergies_dechets"`         // Bioénergies - Déchets (MW) - Production bioénergies réalisée à partir des déchets
	Charbon                  int64  `json:"charbon"`                     // Charbon (MW)
	Consommation             int64  `json:"consommation"`                // Consommation (MW)
	Date                     string `json:"date"`                        // Date
	DateHeure                string `json:"date_heure"`                  // Date - Heure
	DestockageBatterie       string `json:"destockage_batterie"`         // Déstockage batterie (MW)
	EchCommAngleterre        int64  `json:"ech_comm_angleterre"`         // Ech. comm. Angleterre (MW) - Solde des échanges commerciaux entre la France et l'Angleterre. Exportateur si négatif, importateur si positif.
	EchCommEspagne           int64  `json:"ech_comm_espagne"`            // Ech. comm. Espagne (MW) - Solde des échanges commerciaux entre la France et l'Espagne. Exportateur si négatif, importateur si positif.
	EchCommItalie            int64  `json:"ech_comm_italie"`             // Ech. comm. Italie (MW) - Solde des échanges commerciaux entre la France et l'Italie. Exportateur si négatif, importateur si positif.
	EchCommSuisse            int64  `json:"ech_comm_suisse"`             // Ech. comm. Suisse (MW) - Solde des échanges commerciaux entre la France et la Suisse. Exportateur si négatif, importateur si positif.
	EchPhysiques             int64  `json:"ech_physiques"`               // Ech. physiques (MW) - Solde des échanges physiques aux interconnexions avec les autres pays: Exportateur si négatif, importateur si positif.
	Eolien                   int64  `json:"eolien"`                      // Eolien (MW)
	EolienOffshore           string `json:"eolien_offshore"`             // Eolien offshore (MW)
	EolienTerrestre          string `json:"eolien_terrestre"`            // Eolien terrestre (MW)
	Fioul                    int64  `json:"fioul"`                       // Fioul (MW)
	FioulAutres              int64  `json:"fioul_autres"`                // Fioul - Autres (MW)
	FioulCogen               int64  `json:"fioul_cogen"`                 // Fioul - Cogénération (MW) - Production des cogénérations fonctionnant au fioul
	FioulTac                 int64  `json:"fioul_tac"`                   // Fioul - TAC (MW) - Production des Turbines à Combustion fonctionnant au fioul
	Gaz                      int64  `json:"gaz"`                         // Gaz (MW)
	GazAutres                int64  `json:"gaz_autres"`                  // Gaz - Autres (MW) - Autres technologies fonctionnant au gaz
	GazCcg                   int64  `json:"gaz_ccg"`                     // Gaz - CCG (MW) - Production des Cycles Combinés Gaz
	GazTac                   int64  `json:"gaz_tac"`                     // Gaz - TAC (MW) - Production des Turbines à Combustion fonctionnant au gaz
	Heure                    string `json:"heure"`                       // Heure
	Hydraulique              int64  `json:"hydraulique"`                 // Hydraulique (MW)
	HydrauliqueFilEauEclusee int64  `json:"hydraulique_fil_eau_eclusee"` // Hydraulique - Fil de l'eau + éclusée (MW)
	HydrauliqueLacs          int64  `json:"hydraulique_lacs"`            // Hydraulique - Lacs (MW)
	HydrauliqueStepTurbinage int64  `json:"hydraulique_step_turbinage"`  // Hydraulique - STEP turbinage (MW) - Production hydraulique issue des Stations de Transfert d'Energie par Pompage.
	Nature                   string `json:"nature"`                      // Uniquement "Données temps réel" pour ce jeu de données.
	Nucleaire                int64  `json:"nucleaire"`                   // Nucléaire (MW)
	Perimetre                string `json:"perimetre"`                   // Uniquement France pour ce jeu de données.
	Pompage                  int64  `json:"pompage"`                     // Pompage (MW) - Puissance consommée par les pompes dans les Stations de Transfert d'Energie par Pompage (STEP)
	PrevisionJ               int64  `json:"prevision_j"`                 // Prévision J (MW) - Prévision, réalisée le jour même, de la consommation .
	PrevisionJ1              int64  `json:"prevision_j1"`                // Prévision J-1 (MW) - Prévision, réalisée la veille pour le lendemain, de la consommation.
	Solaire                  int64  `json:"solaire"`                     // Solaire (MW)
	StockageBatterie         string `json:"stockage_batterie"`           // Stockage batterie (MW)
	TauxCo2                  int64  `json:"taux_co2"`                    // Taux de CO2 (g/kWh) - Estimation des émissions de carbone générées par la production d'électricité en France.
}

func (client *Eco2mixClient) FetchNationalFinalData(from time.Time, to time.Time, maxResults int) ([]NationalDefinitiveFields, error) {
	params := url.Values{}

	params.Add("dataset", "eco2mix-national-cons-def")
	params.Add("limit", fmt.Sprintf("%d", maxResults))
	params.Add("order_by", "date_heure asc")
	params.Add("where", fmt.Sprintf("taux_co2 is not null AND date_heure>=date'%s' AND date_heure<=date'%s'", from.Format("2006-01-02"), to.Format("2006-01-02"))) // Time filter

	queryString := params.Encode()

	resp, err := client.httpClient.Get(fmt.Sprintf(OPENDATASOFT_API_DEFINITIVE_PATH, client.BaseUrl, queryString))
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
	var data NationalDefinitiveRecord
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %s", err)
	}

	var fields []NationalDefinitiveFields
	fields = data.Results

	return fields, nil
}
