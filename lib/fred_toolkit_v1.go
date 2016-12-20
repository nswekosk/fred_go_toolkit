package lib

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type FredType struct {
	Start        string        `json:"realtime_start" xml:"realtime_start"`
	End          string        `json:"realtime_end" xml:"realtime_end"`
	ObsStart     string        `json:"observation_start" xml:"observation_start"`
	ObsEnd       string        `json:"observation_end" xml:"observation_end"`
	Units        string        `json:"units" xml:"units"`
	OutputType   int           `json:"output_type" xml:"output_type"`
	FileType     string        `json:"file_type" xml:"file_type"`
	OrderBy      string        `json:"order_by" xml:"order_by"`
	SortOrder    string        `json:"sort_order" xml:"sort_order"`
	Count        int           `json:"count" xml:"count"`
	Offset       int           `json:"offset" xml:"offset"`
	Limit        int           `json:"limit" xml:"limit"`
	Categories   []Category    `json:"categories"`
	Release      []Release     `json:"releases" xml:"releases"`
	Seriess      []Series      `json:"seriess" xml:"seriess"`
	Observations []Observation `json:"observations" xml:"observations"`
	VintageDates []string      `json:"vintage_dates" xml:"vintage_dates"`
	Tags         []Tag         `json:"tag" xml:"tag"`
	Sources      []Source      `json:"sources" xml:"sources"`
	ReleaseDates []ReleaseDate `json:"release_dates" xml:"release_dates"`
}

type FredClient struct {
	aPIKEY     string
	fileType   string
	requestURL string
	logFile    *os.File
}

/********************************
 ** CreateFredClient
 **
 ** Creates an instance of a
 ** FRED client.
 ********************************/
func CreateFredClient(APIKey string, FileType ...string) (*FredClient, error) {

	if sameStr(APIKey, "") {
		return nil, errors.New(errorNoAPIKey)
	}

	f, err := os.OpenFile("fred.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file: %v", err.Error())
		return nil, err
	}
	log.SetOutput(f)

	fileType := ""
	if len(FileType) != 0 {
		fileType = FileType[0]
	}

	return &FredClient{
		aPIKEY:     APIKey,
		fileType:   fileType,
		requestURL: apiURL,
		logFile:    f,
	}, nil
}

/********************************
 ** UpdateAPIKEY
 **
 ** Updates the API KEY for the
 ** client.
 ********************************/
func (f *FredClient) UpdateAPIKEY(APIKey string) error {

	if APIKey == "" || len(APIKey) != 32 {
		f.log(errorInvalidAPIKey)
		return errors.New(errorInvalidAPIKey)
	}

	f.aPIKEY = APIKey

	url := strings.Split(f.requestURL, "?")

	f.requestURL = url[0] + "?api_key=" + APIKey

	return nil
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

	url := f.formatUrl(f.requestURL, params, paramType)

	resp, err := http.Get(url)

	if err != nil {
		f.log("[callAPI] Error with HTTP Call: " + err.Error())
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
func (f *FredClient) decodeObj(resp *http.Response, obj *FredType) (*FredType, error) {
	var err error

	switch f.fileType {
	case FileTypeJSON:
		err = json.NewDecoder(resp.Body).Decode(obj)

		if err != nil {
			f.log("[decodeObj] JSON ERROR: " + err.Error())
			return nil, errors.New(errorLibraryFail)
		}
	case FileTypeXML:
		err = xml.NewDecoder(resp.Body).Decode(obj)

		if err != nil {
			f.log("[decodeObj] XML ERROR: " + err.Error())
			return nil, errors.New(errorLibraryFail)
		}
	default:
		err = xml.NewDecoder(resp.Body).Decode(obj)

		if err != nil {
			f.log("[decodeObj] DEFAULT ERROR: " + err.Error())
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
func (f *FredClient) operate(params map[string]interface{}, paramType string) (*FredType, error) {
	if err := f.validateAPIKEY(); err != nil {
		return nil, err
	}

	resp, err := f.callAPI(params, paramType)

	if err != nil {
		f.log("[operate] callAPI Error " + err.Error())
		return nil, err
	}

	obj := &FredType{}

	obj, err = f.decodeObj(resp, obj)

	if err != nil {
		fmt.Printf("[operate] decodeObj Error " + err.Error())
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
func (f *FredClient) formatUrl(url string, params map[string]interface{}, paramType string) string {

	url += paramsLookup[paramType][paramLookupExt].(string)
	firstParam := true

	if len(params) != 0 {
		for paramKey, paramVal := range params {
			if !sameStr(paramKey, "") || !sameStr(paramVal.(string), "") {
				paramOp := "&"
				for _, param := range paramsLookup[paramType][paramLookupParams].([]string) {
					if sameStr(paramKey, param) {
						if firstParam {
							paramOp = "?"
						}

						val := ""
						kind := reflect.TypeOf(paramVal).Kind()

						switch kind {
						case reflect.String:
							val = paramVal.(string)
							val = strings.Replace(val, " ", "+", -1)
							break
						case reflect.Int:
							val = strconv.Itoa(paramVal.(int))
							break
						case reflect.Bool:
							val = strconv.FormatBool(paramVal.(bool))
							break
						}

						url += (paramOp + paramKey + "=" + val)
					}
				}
			}
		}
		url += "&"
	} else {
		url += "/?"
	}

	fileType := ""
	if len(f.fileType) != 0 {
		fileType = f.fileType
	}

	url += "api_key=" + f.aPIKEY + "&file_type=" + fileType

	return url
}

func sameStr(str1 string, str2 string) bool {
	if strings.Compare(str1, str2) == 0 {
		return true
	}
	return false
}

func (f *FredClient) log(logMes string) {
	log.Println(logMes)
}

func (f *FredClient) logError(method string, err error) {
	log.Fatalf("METHOD: %v ERROR: %v", method, err.Error())
	f.logFile.Close()
}
