package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"sort"

	"github.com/damirsharip/contentquo-test/internal/entity"
)

const link = "https://catfact.ninja/breeds"

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) SortBreeds() error {
	response, err := http.Get(link)
	if err != nil {
		return errors.New("error while getting JSON")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.New("error whil reading JSON")
	}

	var cat entity.CatResponse
	err = json.Unmarshal(body, &cat)
	if err != nil {
		return errors.New("error while unmarshalling JSON")
	}

	mp := make(map[string][]string)

	for _, data := range cat.Data {
		mp[data.Country] = append(mp[data.Country], data.Breed)
	}

	for country, breeds := range mp {
		sort.Slice(breeds, func(i, j int) bool {
			return len(breeds[i]) < len(breeds[j])
		})
		mp[country] = breeds
	}

	file, err := os.Create("out.json")
	if err != nil {
		return errors.New("error while create out.json file")
	}

	err = json.NewEncoder(file).Encode(mp)
	if err != nil {
		return errors.New("Error encoding and writing to file:")
	}

	return nil
}
