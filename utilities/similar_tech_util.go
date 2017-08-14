package utilities

import (
	"net/http"
	"io/ioutil"
	"errors"
	"strconv"
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/enkhalifapro/pgen/models"
)

type SimilarTechUtil struct {
}

func callAPI(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// handle 4xx HTTP codes
	if res.StatusCode >= 500 {
		return nil, errors.New("We were unable to complete your request. The following information was supplied \n\n(Request error [status " + strconv.Itoa(res.StatusCode) + "])")
	}
	// handle 5xx HTTP codes
	if res.StatusCode >= 400 && res.StatusCode < 500 {
		return nil, errors.New("We were unable to complete your request. The following information was supplied \n\n(Internal error [status " + strconv.Itoa(res.StatusCode) + "])")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s *SimilarTechUtil) GetTechnologies(siteAddress string) (*models.SimilarTechResponse, error) {
	// call info API
	apiKey := viper.GetString("similar_tech.key")
	url := "https://api.similartech.com/v1/site/" + siteAddress + "/technologies/all/pages?userkey=" + apiKey + "&format=json"

	body, err := callAPI(url)
	if err != nil {
		return nil, err
	}

	// extract result info
	var response models.SimilarTechResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
