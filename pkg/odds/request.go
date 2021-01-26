package odds

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var (
	Client = &http.Client{}
)

func MakeRequest(name string, method string, url string, in interface{}, res interface{}) error {
	return makeRequest(name, method, url, in, res)
}

func makeRequest(name string, method string, url string, in interface{}, res interface{}) error {

	var body []byte
	if method == "POST" && in != nil {
		body, _ = json.Marshal(in)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("client make request %s %s %s - %s(%d)\n", name, method, url, resp.Status, resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New(resp.Status)
}
