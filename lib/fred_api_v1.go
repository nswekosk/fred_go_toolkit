package lib

import (
	"errors"
	"net/http"
)

type FredClient struct {
	aPI_KEY    string
	fileType   string
	requestUrl string
	Categories Category
	Releases   Releases
	Series     Series
	Sources    Seriess
	Tags       Tags
}

const (
	requestUrl = "https://api.stlouisfed.org/fred"
)

func CreateClient(ApiKey string) (*FredClient, error) {

	if ApiKey == "" {
		return nil, errors.New("Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account.")
	}

	return &FredClient{
		aPI_KEY:    ApiKey,
		fileType:   "json",
		requestUrl: requestUrl + "api_key",
	}, nil

}

func (f *FredClient) validateAPIKEY() error {
	if f.aPI_KEY == "" {
		return errors.New("Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account.")
	}
	return nil
}

func (f *FredClient) callAPI(params map[string]interface{}, param_type string) (*http.Response, error) {

	url := formatUrl(f.requestUrl, params, param_type)

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("There was an error in processing the query. Please contact the client administrator.")
	}

	return resp, nil
}

func formatUrl(url string, params map[string]interface{}, param_type string) string {

	for key, val := range params {
		for _, param := range paramsLookup[param_type] {
			if key == param {
				url += ("/?" + key + "=" + val.(string))
			}
		}

	}

	return url

}
