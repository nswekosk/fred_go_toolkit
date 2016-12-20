package lib

type Releases struct {
	Start      string    `json:"realtime_start" xml:"realtime_start"`
	End        string    `json:"realtime_end" xml:"realtime_end"`
	OrderBy    string    `json:"order_by" xml:"order_by"`
	SortOrder  string    `json:"sort_order" xml:"sort_order"`
	Count      int       `json:"count" xml:"count"`
	Offset     int       `json:"offset" xml:"offset"`
	Limit      int       `json:"limit" xml:"limit"`
	ReleaseCol []Release `json:"releases" xml:"releases"`
}

type Release struct {
	ID           int    `json:"id" xml:"id"`
	Start        string `json:"realtime_start" xml:"realtime_start"`
	End          string `json:"realtime_end" xml:"realtime_end"`
	Name         string `json:"name" xml:"name"`
	PressRelease bool   `json:"press_release" xml:"press_release"`
	Link         string `json:"link" xml:"link"`
}

type ReleaseDates struct {
	Start           string        `json:"realtime_start" xml:"realtime_start"`
	End             string        `json:"realtime_end" xml:"realtime_end"`
	OrderBy         string        `json:"order_by" xml:"order_by"`
	SortOrder       string        `json:"sort_order" xml:"sort_order"`
	Count           int           `json:"count" xml:"count"`
	Offset          int           `json:"offset" xml:"offset"`
	Limit           int           `json:"limit" xml:"limit"`
	ReleaseDatesCol []ReleaseDate `json:"release_dates" xml:"release_dates"`
}

type ReleaseDate struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Date string `json:"date" xml:"date"`
}

/********************************
 ** GetReleases
 **
 ** Get all releases of economic
 ** data.
 ********************************/
func (f *FredClient) GetReleases(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releases)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetReleasesDates
 **
 ** Get release dates for all
 ** releases of economic data.
 ********************************/
func (f *FredClient) GetReleasesDates(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releasesDates)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetRelease
 **
 ** Get a release of economic data.
 ********************************/
func (f *FredClient) GetRelease(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, release)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetReleaseDates
 **
 ** Get release dates for a release
 ** of economic data.
 ********************************/
func (f *FredClient) GetReleaseDates(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseDates)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetReleaseSeries
 **
 ** Get the series on a release of
 ** economic data.
 ********************************/
func (f *FredClient) GetReleaseSeries(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseSeries)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetReleaseSources
 **
 ** Get the sources for a Release
 ** of economic data.
 ********************************/
func (f *FredClient) GetReleaseSources(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseSources)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetReleaseTags
 **
 ** Get the tags for a release.
 ********************************/
func (f *FredClient) GetReleaseTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseTags)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetReleaseRelatedTags
 **
 ** Get the related tags for a
 ** release.
 ********************************/
func (f *FredClient) GetReleaseRelatedTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseRelatedTags)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetReleaseTables
 **
 ** Get the related tags for a
 ** category.
 ********************************/
/*
func (f *FredClient) GetReleaseTables(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "RELEASE_TABLES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(tags)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return tags, nil

}
*/
