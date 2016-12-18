package lib

import (
	"encoding/json"
	"errors"
)

type Seriess struct {
	Start     string   `json:"realtime_start"`
	End       string   `json:"realtime_end"`
	SeriesCol []Series `json:"seriess"`
}

type Series struct {
	ID                     string `json:"id"`
	Start                  string `json:"realtime_start"`
	End                    string `json:"realtime_end"`
	Title                  string `json:"title"`
	ObsStart               string `json:"observation_start"`
	ObsEnd                 string `json:"observation_end"`
	Frequency              string `json:"annual"`
	FrequencyShort         string `json:"frequency_short"`
	Units                  string `json:"units"`
	UnitsShort             string `json:"units_short"`
	SeasonalAdjustment     string `json:"seasonal_adjustment"`
	SeasonalAdustmentShort string `json:"seasonal_adjustment_short"`
	LastUpdated            string `json:"last_updated"`
	Popularity             int    `json:"popularity"`
	Notes                  string `json:"notes"`
}

type Observations struct {
	Start        string        `json:"realtime_start"`
	End          string        `json:"realtime_end"`
	ObsStart     string        `json:"observation_start"`
	ObsEnd       string        `json:"observation_end"`
	Units        string        `json:"units"`
	OutputType   int           `json:"output_type"`
	FileType     string        `json:"file_type"`
	OrderBy      string        `json:"order_by"`
	SortOrder    string        `json:"sort_order"`
	Count        int           `json:"count"`
	Offset       int           `json:"offset"`
	Limit        int           `json:"limit"`
	Observations []Observation `json:"observations"`
}

type Observation struct {
	Start string `json:"realtime_start"`
	End   string `json:"realtime_end"`
	Date  string `json:"date"`
	Value string `json:"value"`
}

type VintageDates struct {
	Start        string   `json:"realtime_start"`
	End          string   `json:"realtime_end"`
	OrderBy      string   `json:"order_by"`
	SortOrder    string   `json:"sort_order"`
	Count        int      `json:"count"`
	Offset       int      `json:"offset"`
	Limit        int      `json:"limit"`
	VintageDates []string `json:"vintage_dates"`
}

/********************************
 ** GetSeries
 **
 ** Get an economic data series.
 ********************************/
func (f *FredClient) GetSeries(params map[string]interface{}) (*Seriess, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	srs := &Seriess{}

	resp, err := f.callAPI(params, "SERIES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(srs)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return srs, nil

}

/********************************
 ** GetSeriesCategories
 **
 ** Get the categories for an
 ** economic data series.
 ********************************/
func (f *FredClient) GetSeriesCategories(params map[string]interface{}) (*Categories, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	ctgs := &Categories{}

	resp, err := f.callAPI(params, "SERIES_CATEGORIES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(ctgs)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return ctgs, nil

}

/********************************
 ** GetReleaseObservations
 **
 ** Get the observations or data
 ** values for an economic data
 ** series.
 ********************************/
func (f *FredClient) GetSeriesObservations(params map[string]interface{}) (*Observations, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	obs := &Observations{}

	resp, err := f.callAPI(params, "SERIES_OBSERVATIONS")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(obs)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return obs, nil

}

/********************************
 ** GetSeriesRelease
 **
 ** Get the release for an economic
 ** data series.
 ********************************/
func (f *FredClient) GetSeriesRelease(params map[string]interface{}) (*Releases, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	rls := &Releases{}

	resp, err := f.callAPI(params, "SERIES_RELEASE")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(rls)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return rls, nil

}

/********************************
 ** GetSeriesSearch
 **
 ** Get economic data series that
 ** match keywords.
 ********************************/
func (f *FredClient) GetSeriesSearch(params map[string]interface{}) (*Seriess, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	srs := &Seriess{}

	resp, err := f.callAPI(params, "SERIES_SEARCH")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(srs)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return srs, nil

}

/********************************
 ** GetSeriesSearchTags
 **
 ** Get the tags for a series search.
 ********************************/
func (f *FredClient) GetSeriesSearchTags(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "SERIES_SEARCH_TAGS")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(tags)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return tags, nil

}

/********************************
 ** GetSeriesSearchRelatedTags
 **
 ** Get teh related tags for a
 ** series search.
 ********************************/
func (f *FredClient) GetSeriesSearchRelatedTags(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "SERIES_SEARCH_RELATED_TAGS")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(tags)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return tags, nil

}

/********************************
 ** GetSeriesTags
 **
 ** Get the tags for an economic
 ** data series.
 ********************************/
func (f *FredClient) GetSeriesTags(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "SERIES_TAGS")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(tags)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return tags, nil

}

/********************************
 ** GetSeriesUpdates
 **
 ** Get the economic data series
 ** sorted by when observations
 ** were updated on the FRED
 ** server.
 ********************************/
func (f *FredClient) GetSeriesUpdates(params map[string]interface{}) (*Seriess, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	srs := &Seriess{}

	resp, err := f.callAPI(params, "SERIES_UPDATES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(srs)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return srs, nil

}

/********************************
 ** GetSeriesVintageDates
 **
 ** Get the dates in history when
 ** series' data values were
 ** revised or new data values
 ** were released.
 ********************************/
func (f *FredClient) GetSeriesVintageDates(params map[string]interface{}) (*VintageDates, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	vds := &VintageDates{}

	resp, err := f.callAPI(params, "SERIES_VINTAGEDATES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(vds)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return vds, nil

}
