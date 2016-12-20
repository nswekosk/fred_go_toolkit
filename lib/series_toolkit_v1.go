package lib

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

type Observation struct {
	Start string `json:"realtime_start" xml:"realtime_start"`
	End   string `json:"realtime_end" xml:"realtime_end"`
	Date  string `json:"date" xml:"date"`
	Value string `json:"value" xml:"value"`
}

/********************************
 ** GetSeries
 **
 ** Get an economic data series.
 ********************************/
func (f *FredClient) GetSeries(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, series)

	if err != nil {
		f.logError(series, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSeriesCategories
 **
 ** Get the categories for an
 ** economic data series.
 ********************************/
func (f *FredClient) GetSeriesCategories(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesCategories)

	if err != nil {
		f.logError(seriesCategories, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetReleaseObservations
 **
 ** Get the observations or data
 ** values for an economic data
 ** series.
 ********************************/
func (f *FredClient) GetSeriesObservations(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesObservations)

	if err != nil {
		f.logError(seriesObservations, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSeriesRelease
 **
 ** Get the release for an economic
 ** data series.
 ********************************/
func (f *FredClient) GetSeriesRelease(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesRelease)

	if err != nil {
		f.logError(seriesRelease, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSeriesSearch
 **
 ** Get economic data series that
 ** match keywords.
 ********************************/
func (f *FredClient) GetSeriesSearch(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesSearch)

	if err != nil {
		f.logError(seriesSearch, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSeriesSearchTags
 **
 ** Get the tags for a series search.
 ********************************/
func (f *FredClient) GetSeriesSearchTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesSearchTags)

	if err != nil {
		f.logError(seriesSearchTags, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSeriesSearchRelatedTags
 **
 ** Get teh related tags for a
 ** series search.
 ********************************/
func (f *FredClient) GetSeriesSearchRelatedTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesSearchRelatedTags)

	if err != nil {
		f.logError(seriesSearchRelatedTags, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSeriesTags
 **
 ** Get the tags for an economic
 ** data series.
 ********************************/
func (f *FredClient) GetSeriesTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesTags)

	if err != nil {
		f.logError(seriesTags, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSeriesUpdates
 **
 ** Get the economic data series
 ** sorted by when observations
 ** were updated on the FRED
 ** server.
 ********************************/
func (f *FredClient) GetSeriesUpdates(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesUpdates)

	if err != nil {
		f.logError(seriesUpdates, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSeriesVintageDates
 **
 ** Get the dates in history when
 ** series' data values were
 ** revised or new data values
 ** were released.
 ********************************/
func (f *FredClient) GetSeriesVintageDates(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesVintagedates)

	if err != nil {
		f.logError(seriesVintagedates, err)
		return nil, err
	}

	return fc, nil

}
