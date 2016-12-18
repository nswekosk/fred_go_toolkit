package lib

import (
	"encoding/json"
	"errors"
)

type Releases struct {
	Start      string    `realtime_start`
	End        string    `realtime_end`
	OrderBy    string    `order_by`
	SortOrder  string    `sort_order`
	Count      int       `count`
	Offset     int       `offset`
	Limit      int       `limit`
	ReleaseCol []Release `releases`
}

type Release struct {
	ID           int    `id`
	Start        string `realtime_start`
	End          string `realtime_end`
	Name         string `name`
	PressRelease bool   `press_release`
	Link         string `link`
}

type ReleaseDates struct {
	Start           string        `realtime_start`
	End             string        `realtime_end`
	OrderBy         string        `order_by`
	SortOrder       string        `sort_order`
	Count           int           `count`
	Offset          int           `offset`
	Limit           int           `limit`
	ReleaseDatesCol []ReleaseDate `release_dates`
}

type ReleaseDate struct {
	ID   int    `id`
	Name string `name`
	Date string `date`
}

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
