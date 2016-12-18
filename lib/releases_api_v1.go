package lib

import (
	"encoding/json"
	"errors"
)

type Releases struct {
	Start      string    `json:"realtime_start"`
	End        string    `json:"realtime_end"`
	OrderBy    string    `json:"order_by"`
	SortOrder  string    `json:"sort_order"`
	Count      int       `json:"count"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	ReleaseCol []Release `json:"releases"`
}

type Release struct {
	ID           int    `json:"id"`
	Start        string `json:"realtime_start"`
	End          string `json:"realtime_end"`
	Name         string `json:"name"`
	PressRelease bool   `json:"press_release"`
	Link         string `json:"link"`
}

type ReleaseDates struct {
	Start           string        `json:"realtime_start"`
	End             string        `json:"realtime_end"`
	OrderBy         string        `json:"order_by"`
	SortOrder       string        `json:"sort_order"`
	Count           int           `json:"count"`
	Offset          int           `json:"offset"`
	Limit           int           `json:"limit"`
	ReleaseDatesCol []ReleaseDate `json:"release_dates"`
}

type ReleaseDate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

/********************************
 ** GetReleases
 **
 ** Get all releases of economic
 ** data.
 ********************************/
func (f *FredClient) GetReleases(params map[string]interface{}) (*Releases, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	rls := &Releases{}

	resp, err := f.callAPI(params, "RELEASES")

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
 ** GetReleasesDates
 **
 ** Get release dates for all
 ** releases of economic data.
 ********************************/
func (f *FredClient) GetReleasesDates(params map[string]interface{}) (*ReleaseDates, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	rlDts := &ReleaseDates{}

	resp, err := f.callAPI(params, "RELEASES_DATES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(rlDts)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return rlDts, nil

}

/********************************
 ** GetRelease
 **
 ** Get a release of economic data.
 ********************************/
func (f *FredClient) GetRelease(params map[string]interface{}) (*Releases, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	rls := &Releases{}

	resp, err := f.callAPI(params, "RELEASE")

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
 ** GetReleaseDates
 **
 ** Get release dates for a release
 ** of economic data.
 ********************************/
func (f *FredClient) GetReleaseDates(params map[string]interface{}) (*ReleaseDates, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	rlDts := &ReleaseDates{}

	resp, err := f.callAPI(params, "RELEASE_DATES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(rlDts)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return rlDts, nil

}

/********************************
 ** GetReleaseSeries
 **
 ** Get the series on a release of
 ** economic data.
 ********************************/
func (f *FredClient) GetReleaseSeries(params map[string]interface{}) (*Seriess, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	srs := &Seriess{}

	resp, err := f.callAPI(params, "RELEASE_SERIES")

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
 ** GetReleaseSources
 **
 ** Get the sources for a Release
 ** of economic data.
 ********************************/
func (f *FredClient) GetReleaseSources(params map[string]interface{}) (*Sources, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	srcs := &Sources{}

	resp, err := f.callAPI(params, "RELEASE_SOURCES")

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
 ** GetReleaseTags
 **
 ** Get the tags for a release.
 ********************************/
func (f *FredClient) GetReleaseTags(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "RELEASE_SOURCES")

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
 ** GetReleaseRelatedTags
 **
 ** Get the related tags for a
 ** release.
 ********************************/
func (f *FredClient) GetReleaseRelatedTags(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "RELEASE_RELATED_TAGS")

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
