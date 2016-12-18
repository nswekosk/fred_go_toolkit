package lib

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

	srs, err := f.operate(params, "SERIES")

	if err != nil {
		return nil, err
	}

	return (srs.(*Seriess)), nil

}

/********************************
 ** GetSeriesCategories
 **
 ** Get the categories for an
 ** economic data series.
 ********************************/
func (f *FredClient) GetSeriesCategories(params map[string]interface{}) (*Categories, error) {

	srsCts, err := f.operate(params, "SERIES_CATEGORIES")

	if err != nil {
		return nil, err
	}

	return (srsCts.(*Categories)), nil

}

/********************************
 ** GetReleaseObservations
 **
 ** Get the observations or data
 ** values for an economic data
 ** series.
 ********************************/
func (f *FredClient) GetSeriesObservations(params map[string]interface{}) (*Observations, error) {

	srsObs, err := f.operate(params, "SERIES_OBSERVATIONS")

	if err != nil {
		return nil, err
	}

	return (srsObs.(*Observations)), nil

}

/********************************
 ** GetSeriesRelease
 **
 ** Get the release for an economic
 ** data series.
 ********************************/
func (f *FredClient) GetSeriesRelease(params map[string]interface{}) (*Releases, error) {

	srsRls, err := f.operate(params, "SERIES_RELEASE")

	if err != nil {
		return nil, err
	}

	return (srsRls.(*Releases)), nil

}

/********************************
 ** GetSeriesSearch
 **
 ** Get economic data series that
 ** match keywords.
 ********************************/
func (f *FredClient) GetSeriesSearch(params map[string]interface{}) (*Seriess, error) {

	srs, err := f.operate(params, "SERIES_SEARCH")

	if err != nil {
		return nil, err
	}

	return (srs.(*Seriess)), nil

}

/********************************
 ** GetSeriesSearchTags
 **
 ** Get the tags for a series search.
 ********************************/
func (f *FredClient) GetSeriesSearchTags(params map[string]interface{}) (*Tags, error) {

	tags, err := f.operate(params, "SERIES_SEARCH_TAGS")

	if err != nil {
		return nil, err
	}

	return (tags.(*Tags)), nil

}

/********************************
 ** GetSeriesSearchRelatedTags
 **
 ** Get teh related tags for a
 ** series search.
 ********************************/
func (f *FredClient) GetSeriesSearchRelatedTags(params map[string]interface{}) (*Tags, error) {

	tags, err := f.operate(params, "SERIES_SEARCH_RELATED_TAGS")

	if err != nil {
		return nil, err
	}

	return (tags.(*Tags)), nil

}

/********************************
 ** GetSeriesTags
 **
 ** Get the tags for an economic
 ** data series.
 ********************************/
func (f *FredClient) GetSeriesTags(params map[string]interface{}) (*Tags, error) {

	tags, err := f.operate(params, "SERIES_TAGS")

	if err != nil {
		return nil, err
	}

	return (tags.(*Tags)), nil

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

	srs, err := f.operate(params, "SERIES_UPDATES")

	if err != nil {
		return nil, err
	}

	return (srs.(*Seriess)), nil

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

	vds, err := f.operate(params, "SERIES_VINTAGEDATES")

	if err != nil {
		return nil, err
	}

	return (vds.(*VintageDates)), nil

}
