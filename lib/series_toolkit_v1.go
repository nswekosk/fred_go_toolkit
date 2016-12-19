package lib

type Seriess struct {
	Start     string   `json:"realtime_start" xml:"realtime_start"`
	End       string   `json:"realtime_end" xml:"realtime_end"`
	SeriesCol []Series `json:"seriess" xml:"seriess"`
}

type Series struct {
	ID                     string `json:"id" xml:"id"`
	Start                  string `json:"realtime_start" xml:"realtime_start"`
	End                    string `json:"realtime_end" xml:"realtime_end"`
	Title                  string `json:"title" xml:"title"`
	ObsStart               string `json:"observation_start" xml:"observation_start"`
	ObsEnd                 string `json:"observation_end" xml:"observation_end"`
	Frequency              string `json:"annual" xml:"annual"`
	FrequencyShort         string `json:"frequency_short" xml:"frequency_short"`
	Units                  string `json:"units" xml:"units"`
	UnitsShort             string `json:"units_short" xml:"units_short"`
	SeasonalAdjustment     string `json:"seasonal_adjustment" xml:"seasonal_adjustment"`
	SeasonalAdustmentShort string `json:"seasonal_adjustment_short" xml:"seasonal_adjustment_short"`
	LastUpdated            string `json:"last_updated" xml:"last_updated"`
	Popularity             int    `json:"popularity" xml:"popularity"`
	Notes                  string `json:"notes" xml:"notes"`
}

type Observations struct {
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
	Observations []Observation `json:"observations" xml:"observations"`
}

type Observation struct {
	Start string `json:"realtime_start" xml:"realtime_start"`
	End   string `json:"realtime_end" xml:"realtime_end"`
	Date  string `json:"date" xml:"date"`
	Value string `json:"value" xml:"value"`
}

type VintageDates struct {
	Start        string   `json:"realtime_start" xml:"realtime_start"`
	End          string   `json:"realtime_end" xml:"realtime_end"`
	OrderBy      string   `json:"order_by" xml:"order_by"`
	SortOrder    string   `json:"sort_order" xml:"sort_order"`
	Count        int      `json:"count" xml:"count"`
	Offset       int      `json:"offset" xml:"offset"`
	Limit        int      `json:"limit" xml:"limit"`
	VintageDates []string `json:"vintage_dates" xml:"vintage_dates"`
}

/********************************
 ** GetSeries
 **
 ** Get an economic data series.
 ********************************/
func (f *FredClient) GetSeries(params map[string]interface{}) (*Seriess, error) {

	srs, err := f.operate(params, series)

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

	srsCts, err := f.operate(params, seriesCategories)

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

	srsObs, err := f.operate(params, seriesObservations)

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

	srsRls, err := f.operate(params, seriesRelease)

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

	srs, err := f.operate(params, seriesSearch)

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

	tags, err := f.operate(params, seriesSearchTags)

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

	tags, err := f.operate(params, seriesSearchRelatedTags)

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

	tags, err := f.operate(params, seriesTags)

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

	srs, err := f.operate(params, seriesUpdates)

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

	vds, err := f.operate(params, seriesVintagedates)

	if err != nil {
		return nil, err
	}

	return (vds.(*VintageDates)), nil

}
