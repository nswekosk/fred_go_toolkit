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

const (
	requestUrl = "https://api.stlouisfed.org/fred"
)

func CreateClient(ApiKey string) (*FredClient, error) {

	if sameStr(ApiKey, "") {
		return nil, errors.New("Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account.")
	}

	return &FredClient{
		aPI_KEY:    ApiKey,
		fileType:   "json",
		requestUrl: requestUrl + "?api_key=" + ApiKey,
	}, nil
}

func (f *FredClient) UpdateAPIKEY(ApiKey string) {

	f.aPI_KEY = ApiKey

	url := strings.Split(f.requestUrl, "?")

	f.requestUrl = url[0] + "?api_key=" + ApiKey

}

func (f *FredClient) validateAPIKEY() error {
	if sameStr(f.aPI_KEY, "") {
		return errors.New("Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account.")
	}
	return nil
}

func (f *FredClient) callAPI(params map[string]interface{}, param_type string) (*http.Response, error) {

	url := formatUrl(f.requestUrl, params, param_type)

	if sameStr(url, f.requestUrl) {
		return nil, errors.New("No parameters were added. Please update you parameter input.")
	}

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("There was an error in processing the query. Please contact the client administrator.")
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
