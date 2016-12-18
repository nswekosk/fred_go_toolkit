package lib

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

type FredInterface interface{}

type FredClient struct {
	aPIKEY     string
	fileType   string
	requestURL string
}

/********************************
 ** CreateClient
 **
 ** Creates an instance of a
 ** FRED client.
 ********************************/
func CreateClient(APIKey string, FileType ...string) (*FredClient, error) {

	if sameStr(APIKey, "") {
		return nil, errors.New(errorNoAPIKey)
	}

	return &FredClient{
		aPIKEY:     APIKey,
		fileType:   FileType[0],
		requestURL: apiURL + "?aPIKEY=" + APIKey,
	}, nil
}

/********************************
 ** UpdateAPIKEY
 **
 ** Updates the API KEY for the
 ** client.
 ********************************/
func (f *FredClient) UpdateAPIKEY(APIKey string) {

	f.aPIKEY = APIKey

	url := strings.Split(f.requestURL, "?")

	f.requestURL = url[0] + "?aPIKEY=" + APIKey

}

/********************************
 ** validateAPIKEY
 **
 ** Validates that an APIKEY exists.
 ********************************/
func (f *FredClient) validateAPIKEY() error {
	if sameStr(f.aPIKEY, "") {
		return errors.New(errorNoAPIKey)
	}
	return nil
}

/********************************
 ** callAPI
 **
 ** Creates the url and makes a
 ** GET request to the API.
 ********************************/
func (f *FredClient) callAPI(params map[string]interface{}, paramType string) (*http.Response, error) {

	url := formatUrl(f.requestURL, params, paramType)

	if sameStr(url, f.requestURL) {
		return nil, errors.New(errorNoParameters)
	}

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return resp, nil
}

/********************************
 ** decodeObj
 **
 ** Decodes the object in the
 ** format specified by ther user.
 ********************************/
func (f *FredClient) decodeObj(resp *http.Response, obj FredInterface) (FredInterface, error) {
	var err error

	switch f.fileType {
	case FileTypeJSON:
		err = json.NewDecoder(resp.Body).Decode(obj)

		if err != nil {
			return nil, errors.New(errorLibraryFail)
		}
	case FileTypeXML:
		err = xml.NewDecoder(resp.Body).Decode(obj)

		if err != nil {
			return nil, errors.New(errorLibraryFail)
		}
	default:
		err = xml.NewDecoder(resp.Body).Decode(obj)

		if err != nil {
			return nil, errors.New(errorLibraryFail)
		}

	}

	return obj, nil

}

/********************************
 ** operate
 **
 ** Runs the operation based
 ** parameter type.
 ********************************/
func (f *FredClient) operate(params map[string]interface{}, paramType string) (FredInterface, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	resp, err := f.callAPI(params, paramType)

	if err != nil {
		return nil, err
	}

	var obj FredInterface

	obj, err = f.decodeObj(resp, obj)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

/********************************
 ** formatUrl
 **
 ** Formats the url per the API
 ** specifications.
 ********************************/
func formatUrl(url string, params map[string]interface{}, paramType string) string {

	url += paramsLookup[paramType][paramLookupExt].(string)
	firstParam := true

	for paramKey, paramVal := range params {
		paramOp := "&"
		for _, param := range paramsLookup[paramType][paramLookupParams].([]string) {
			if sameStr(paramKey, param) {
				if firstParam {
					paramOp = "?"

				}
				url += (paramOp + paramKey + "=" + paramVal.(string))
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
