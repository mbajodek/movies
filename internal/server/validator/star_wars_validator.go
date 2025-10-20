package validator

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

var peopleUrl = "https://swapi.dev/api/people/"

type CharacterSearchResult struct {
	Count int `json:"count"`
}

type StarWarsValidator struct {
	restyClient *resty.Client
}

func NewStarWarsValidator() *StarWarsValidator {
	return &StarWarsValidator{
		restyClient: resty.New(),
	}
}

func (validator StarWarsValidator) CharacterExistsInMovie(name string) (bool, error) {
	res, err := validator.restyClient.R().
		SetQueryParam("search", name).
		Get(peopleUrl)

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