package lib

import (
	"encoding/json"
	"errors"
)

type Seriess struct {
	Start     string   `realtime_start`
	End       string   `realtime_end`
	SeriesCol []Series `seriess`
}

type Series struct {
	ID                     string `id`
	Start                  string `realtime_start`
	End                    string `realtime_end`
	Title                  string `title`
	ObsStart               string `observation_start`
	ObsEnd                 string `observation_end`
	Frequency              string `annual`
	FrequencyShort         string `frequency_short`
	Units                  string `units`
	UnitsShort             string `units_short`
	SeasonalAdjustment     string `seasonal_adjustment`
	SeasonalAdustmentShort string `seasonal_adjustment_short`
	LastUpdated            string `last_updated`
	Popularity             int    `popularity`
	Notes                  string `notes`
}

type Observations struct {
	Start        string        `realtime_start`
	End          string        `realtime_end`
	ObsStart     string        `observation_start`
	ObsEnd       string        `observation_end`
	Units        string        `units`
	OutputType   int           `output_type`
	FileType     string        `file_type`
	OrderBy      string        `order_by`
	SortOrder    string        `sort_order`
	Count        int           `count`
	Offset       int           `offset`
	Limit        int           `limit`
	Observations []Observation `observations`
}

type Observation struct {
	Start string `realtime_start`
	End   string `realtime_end`
	Date  string `date`
	Value string `value`
}

type VintageDates struct {
	Start        string   `realtime_start`
	End          string   `realtime_end`
	OrderBy      string   `order_by`
	SortOrder    string   `sort_order`
	Count        int      `count`
	Offset       int      `offset`
	Limit        int      `limit`
	VintageDates []string `vintage_dates`
}

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
