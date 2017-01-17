package fred_go_toolkit

import (
	"errors"
	"io/ioutil"
)

// Release is a single instance of a release of FRED economic data.
type Release struct {
	ID           int    `json:"id" xml:"id,attr"`
	Start        string `json:"realtime_start" xml:"realtime_start,attr"`
	End          string `json:"realtime_end" xml:"realtime_end,attr"`
	Name         string `json:"name" xml:"name,attr"`
	PressRelease bool   `json:"press_release" xml:"press_release,attr"`
	Link         string `json:"link" xml:"link,attr"`
}

// ReleaseDate is a single instance of a release date of FRED economic data.
type ReleaseDate struct {
	ID   int    `json:"release_id" xml:"release_id,attr"`
	Name string `json:"release_name" xml:"release_name,attr"`
	Date string `json:"date" xml:",chardata"`
}

// GetReleases will get all releases of economic data.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/releases.html
func (f *FredClient) GetReleases(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releasesParam)

	if err != nil {
		f.logError(releasesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetReleasesDates will get release dates for all releases of economic data.
// Note that release dates are published by data sources and do not necessarily represent when data will be available on the FRED or ALFRED websites.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/releases_dates.html
func (f *FredClient) GetReleasesDates(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseDatesParam)

	if err != nil {
		f.logError(releaseDatesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetRelease will get a release of economic data.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/release.html
func (f *FredClient) GetRelease(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseParam)

	if err != nil {
		f.logError(releaseParam, err)
		return nil, err
	}

	return fc, nil

}

// GetReleaseDates will get release dates for a release of economic data.
// Note that release dates are published by data sources and do not necessarily represent when data will be available on the FRED or ALFRED websites.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/release_dates.html
func (f *FredClient) GetReleaseDates(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseDatesParam)

	if err != nil {
		f.logError(releaseDatesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetReleaseSeries will get the series on a release of economic data.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/release_series.html
func (f *FredClient) GetReleaseSeries(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseSeriesParam)

	if err != nil {
		f.logError(releaseSeriesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetReleaseSources will get the sources for a Release of economic data.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/release_sources.html
func (f *FredClient) GetReleaseSources(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseSourcesParam)

	if err != nil {
		f.logError(releaseSourcesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetReleaseTags will get the FRED tags for a release.
// Optionally, filter results by tag name, tag group, or search.
// Series are assigned tags and releases.
// Indirectly through series, it is possible to get the tags for a release.
// See the related request GetReleaseRelatedTags.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/release_tags.html
func (f *FredClient) GetReleaseTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseTagsParam)

	if err != nil {
		f.logError(releaseTagsParam, err)
		return nil, err
	}

	return fc, nil

}

// GetReleaseRelatedTags will get the related FRED tags for one or more FRED tags within a release.
// Optionally, filter results by tag group or search.
//
// FRED tags are attributes assigned to series.
// For this request, related FRED tags are the tags assigned to series that match all tags in the tag_names parameter, no tags in the exclude_tag_names parameter, and the release set by the release_id parameter.
// See the related request GetReleaseTags.
//
// Series are assigned tags and releases.
// Indirectly through series, it is possible to get the tags for a release.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/release_related_tags.html
func (f *FredClient) GetReleaseRelatedTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, releaseRelatedTagsParam)

	if err != nil {
		f.logError(releaseRelatedTagsParam, err)
		return nil, err
	}

	return fc, nil

}

// GetReleaseTables will get release table trees for a given release.
// You can go directly to the tree structure by passing the appropriate ElementId.
// You may also use a drill-down approach to start at the root (top most) element by leaving the element_id off.
//
// Note that release dates are published by data sources and do not necessarily represent when data will be available on the FRED or ALFRED websites.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/release_tables.html
func (f *FredClient) GetReleaseTables(params map[string]interface{}) (string, error) {
	if err := f.validateAPIKEY(); err != nil {

		return "", err

	}

	resp, err := f.callAPI(params, "RELEASE_TABLES")

	if err != nil {
		return "", err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		f.log("[GetReleaseTables] READ ERROR: " + err.Error())
		return "", errors.New(errorLibraryFail)
	}

	return string(res), nil

}
