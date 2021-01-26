package odds

import (
	"fmt"
	"net/http"

	"github.com/Ariesfall/simple-odds-api/pkg/data"
)

const (
	Sport  = "upcoming"
	Region = "uk"
	Market = "h2h"
)

var (
	Endpoint = "https://api.the-odds-api.com"
	ApiKey   = ""
)

// OddsResp is the response of odds request
type OddsResp struct {
	Success bool        `json:"success"`
	Data    []*OddsData `json:"data"`
}

// OddsData is the detail of odd result data
type OddsData struct {
	SportKey     string   `json:"sport_key"`
	SportNice    string   `json:"sport_nice"`
	Teams        []string `json:"teams"`
	CommenceTime int      `json:"commence_time"`
	HomeTeam     string   `json:"home_team"`
	Sites        []*Site  `json:"sites"`
	SitesCnt     int      `json:"sites_count"`
}

type Site struct {
	SiteKey    string `json:"site_key"`
	SiteNice   string `json:"site_nice"`
	LastUpdate int    `json:"last_update"`
	Odds       *Odds  `json:"odds"`
}

type Odds struct {
	H2h []float32 `json:"h2h"` // team A, team B, draw
}

// GetOdds send request to odds to get the data of matchs and odds
// sport: Sport key obtained from calling the /sports, upcoming is always valid
// region: Valid regions are au (Australia), uk (United Kingdom), eu (Europe) and us (United States)
// mkt: Optional - Determines which odds market is returned. Defaults to h2h (head to head / moneyline).
func GetOdds(sport string) (*OddsResp, error) {
	if sport == "" {
		sport = Sport
	}

	url := fmt.Sprintf("%s/v3/odds?apiKey=%s&sport=%s&region=%s&mkt=%s", Endpoint, ApiKey, sport, Region, Market)
	res := &OddsResp{}

	err := makeRequest("getOdds", http.MethodGet, url, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Sports is the response of sport request
type SportsResp struct {
	Success bool           `json:"success"`
	Data    []*data.Sports `json:"data"`
}

// GetSports send request to odds to get the data of sports
func GetSports() (*SportsResp, error) {
	url := Endpoint + "/v3/sports?apiKey=" + ApiKey
	res := &SportsResp{}

	err := makeRequest("getSport", http.MethodGet, url, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
