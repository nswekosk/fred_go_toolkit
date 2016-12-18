package lib

import "errors"

const requestUrl = "https://api.stlouisfed.org/fred"

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

	if ApiKey == "" {
		return nil, errors.New("Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account.")
	}

	return &FredClient{
		aPI_KEY:    ApiKey,
		fileType:   "json",
		requestUrl: requestUrl + "api_key",
	}, nil

}

func (f *FredClient) CallAPI() error {

	if err := f.validateAPIKEY(); err != nil {

		return err

	}
	return nil
}

func (f *FredClient) validateAPIKEY() error {
	if f.aPI_KEY == "" {
		return errors.New("Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account.")
	}
	return nil
}
