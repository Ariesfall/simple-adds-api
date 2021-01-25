package odds

import (
	"fmt"
	"net/http"

	"github.com/Ariesfall/simple-odds-api/pkg/request"
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
	SportKey  string   `json:"sport_key"`
	SportNice string   `json:"sport_nice"`
	Teams     []string `json:"teams"`
	CommTime  int      `json:"commence_time"`
	HomeTeam  string   `json:"home_team"`
	Sites     []*Site  `json:"sites"`
	SitesCnt  int      `json:"sites_count"`
}

type Site struct {
	SiteKey    string `json:"site_key"`
	SiteNice   string `json:"site_nice"`
	LastUpdate int    `json:"last_update"`
	Odds       *Odds  `json:"odds"`
}

type Odds struct {
	H2h []float64 `json:"h2h"`
}

func GetOdds(sport, region, mkt string) (*OddsResp, error) {
	url := fmt.Sprintf("%s/v3/sports/?apiKey=%s&sport=%s&region=%s&mkt=%s", Endpoint, ApiKey, sport, region, mkt)
	res := &OddsResp{}

	err := request.MakeRequest("getOdds", http.MethodGet, url, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
