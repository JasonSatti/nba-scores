package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Games struct {
	Results int `json:"results"`
	Info    []struct {
		Arena struct {
			Name    string `json:"name"`
			City    string `json:"city"`
			State   string `json:"state"`
			Country string `json:"country"`
		} `json:"arena"`
		Team struct {
			Visitor struct {
				Name     string `json:"name"`
				Nickname string `json:"nickname"`
				Code     string `json:"code"`
				Logo     string `json:"logo"`
			} `json:"visitors"`
			Home struct {
				Name     string `json:"name"`
				Nickname string `json:"nickname"`
				Code     string `json:"code"`
				Logo     string `json:"logo"`
			} `json:"home"`
		} `json:"teams"`
		Scores struct {
			Visitor struct {
				Linescore []string `json:"linescore"`
				Points    int      `json:"points"`
			} `json:"visitors"`
			Home struct {
				Linescore []string `json:"linescore"`
				Points    int      `json:"points"`
			} `json:"home"`
		} `json:"scores"`
		Officials   []string `json:"officials"`
		TimesTied   int      `json:"timesTied"`
		LeadChanges int      `json:"leadChanges"`
	} `json:"response"`
}

func GetGames(date string, key string) (games Games, err error) {
	url := fmt.Sprintf("https://api-nba-v1.p.rapidapi.com/games?date=%s", date)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return games, errors.Wrap(err, fmt.Sprintf("error creating new http call with method: %s to url: %s", req.Method, req.URL))
	}
	req.Header.Add("X-RapidAPI-Key", key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return games, errors.Wrap(err, fmt.Sprintf("error sending the http call with method: %s to url: %s", req.Method, req.URL))
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			errors.Wrap(err, "error attempting to close the response body")
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return games, errors.Wrap(err, "error reading the response body")
	}

	json.Unmarshal(body, &games)
	if err != nil {
		return games, errors.Wrap(err, "error decoding the response body")
	}

	return games, nil

}
