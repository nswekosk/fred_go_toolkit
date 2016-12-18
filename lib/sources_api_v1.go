package lib

import (
	"encoding/json"
	"errors"
)

type Sources struct {
	Start      string   `json:"realtime_start"`
	End        string   `json:"realtime_end"`
	OrderBy    string   `json:"order_by"`
	SortOrder  string   `json:"sort_order"`
	Count      int      `json:"count"`
	Offset     int      `json:"offset"`
	Limit      int      `json:"limit"`
	SourcesCol []Source `json:"sources"`
}

type Source struct {
	ID    int    `json:"id"`
	Start string `json:"realtime_start"`
	End   string `json:"realtime_end"`
	Name  string `json:"name"`
	Link  string `json:"link"`
}

/********************************
 ** GetSources
 **
 ** Get all sources of economic
 ** data.
 ********************************/
func (f *FredClient) GetSources(params map[string]interface{}) (*Sources, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	srcs := &Sources{}

	resp, err := f.callAPI(params, "SOURCES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(srcs)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return srcs, nil

}

/********************************
 ** GetSource
 **
 ** Get a source of economic data.
 ********************************/
func (f *FredClient) GetSource(params map[string]interface{}) (*Sources, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	srcs := &Sources{}

	resp, err := f.callAPI(params, "SOURCE")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(srcs)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return srcs, nil

}

/********************************
 ** GetSourceReleases
 **
 ** Get the releases for a source.
 ********************************/
func (f *FredClient) GetSourceReleases(params map[string]interface{}) (*Releases, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	rls := &Releases{}

	resp, err := f.callAPI(params, "SOURCE_RELEASES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(rls)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return rls, nil

}
