package lib

import (
	"errors"
	"net/http"
	"strings"
)

const (
	apiUrl            = "https://api.stlouisfed.org/fred"
	errorNoApiKey     = "Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account."
	errorNoParameters = "No parameters were added. Please update you parameter input."
	errorLibraryFail  = "There was an error in processing the query. Please contact the client administrator."
)

type FredClient struct {
	aPI_KEY    string
	fileType   string
	requestUrl string
}

/********************************
 ** CreateClient
 **
 ** Creates an instance of a
 ** FRED client.
 ********************************/
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

/********************************
 ** UpdateAPIKEY
 **
 ** Updates the API KEY for the
 ** client.
 ********************************/
func (f *FredClient) UpdateAPIKEY(ApiKey string) {

	f.aPI_KEY = ApiKey

	url := strings.Split(f.requestUrl, "?")

	f.requestUrl = url[0] + "?api_key=" + ApiKey

}

/********************************
 ** validateAPIKEY
 **
 ** Validates that an APIKEY exists.
 ********************************/
func (f *FredClient) validateAPIKEY() error {
	if sameStr(f.aPI_KEY, "") {
		return errors.New(errorNoApiKey)
	}
	return nil
}

/********************************
 ** callAPI
 **
 ** Creates the url and makes a
 ** GET request to the API.
 ********************************/
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

/********************************
 ** formatUrl
 **
 ** Formats the url per the API
 ** specifications.
 ********************************/
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
