package lib

type Release struct {
	ID           int    `json:"id" xml:"id,attr"`
	Start        string `json:"realtime_start" xml:"realtime_start,attr"`
	End          string `json:"realtime_end" xml:"realtime_end,attr"`
	Name         string `json:"name" xml:"name,attr"`
	PressRelease bool   `json:"press_release" xml:"press_release,attr"`
	Link         string `json:"link" xml:"link,attr"`
}

type ReleaseDate struct {
	ID   int    `json:"release_id" xml:"release_id,attr"`
	Name string `json:"release_name" xml:"release_name,attr"`
	Date string `json:"date" xml:"date,attr"`
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
		f.logError(releases, err)
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
		f.logError(releasesDates, err)
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
		f.logError(release, err)
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
		f.logError(releaseDates, err)
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
		f.logError(releaseSeries, err)
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
		f.logError(releaseSources, err)
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
		f.logError(releaseTags, err)
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
		f.logError(releaseRelatedTags, err)
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
