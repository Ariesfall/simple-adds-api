package odds

import (
	"net/http"

	"github.com/Ariesfall/simple-odds-api/pkg/request"
)

// Sports is the response of sport request
type SportsResp struct {
	Success bool      `json:"success"`
	Data    []*Sports `json:"data"`
}

// Sports is the detail of sport result data
type Sports struct {
	Key    string `json:"key"`
	Active bool   `json:"active"`
	Group  string `json:"group"`
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

func GetSports() (*SportsResp, error) {
	url := Endpoint + "/v3/sports/?apiKey=" + ApiKey
	res := &SportsResp{}

	err := request.MakeRequest("getSport", http.MethodGet, url, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
