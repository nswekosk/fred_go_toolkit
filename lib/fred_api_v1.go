package lib

import (
	"errors"
	"net/http"
	"strings"
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

func CreateClient(ApiKey string) (*FredClient, error) {

	if sameStr(ApiKey, "") {
		return nil, errors.New(errorNoApiKey)
	}

	return &FredClient{
		aPI_KEY:    ApiKey,
		fileType:   "json",
		requestUrl: apiUrl + "?api_key=" + ApiKey,
	}, nil
}

func (f *FredClient) UpdateAPIKEY(ApiKey string) {

	f.aPI_KEY = ApiKey

	url := strings.Split(f.requestUrl, "?")

	f.requestUrl = url[0] + "?api_key=" + ApiKey

}

func (f *FredClient) validateAPIKEY() error {
	if sameStr(f.aPI_KEY, "") {
		return errors.New(errorNoApiKey)
	}
	return nil
}

func (f *FredClient) callAPI(params map[string]interface{}, param_type string) (*http.Response, error) {

	url := formatUrl(f.requestUrl, params, param_type)

	if sameStr(url, f.requestUrl) {
		return nil, errors.New(errorNoParameters)
	}

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return resp, nil
}

func formatUrl(url string, params map[string]interface{}, param_type string) string {

	url += paramsLookup[param_type]["extension"].(string)

	for parameter, paramVal := range params {
		for _, param := range paramsLookup[param_type]["params"].([]string) {
			if sameStr(parameter, param) {
				url += ("/?" + parameter + "=" + paramVal.(string))
			}
		}
	}
	return url
}

func sameStr(str1 string, str2 string) bool {
	if strings.Compare(str1, str2) == 0 {
		return true
	}
	return false
}
