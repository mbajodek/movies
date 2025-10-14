package server

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type CharacterSearchResult struct {
	Count int `json:"count"`
}

type StarWarsValidator struct {
	restyClient *resty.Client
}

func NewStarWarsValidator(rc *resty.Client) *StarWarsValidator {
	return &StarWarsValidator{restyClient: rc}
}

func (validator StarWarsValidator) CharacterExistsInMovie(name string) (bool, error) {
	url := strings.ReplaceAll(fmt.Sprintf("https://swapi.dev/api/people/?search=%s", name), " ", "+")
	res, err := validator.restyClient.R().Get(url)

	if err != nil {
		fmt.Println("swapi error:" , err)
		return false, err
	}

	c := CharacterSearchResult{}
	err = json.Unmarshal(res.Body(), &c)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return false, err
	}

	if c.Count == 0 {
		return false, nil
	}

	return true, nil
}