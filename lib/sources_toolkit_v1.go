package lib

import (
	. "github.com/nswekosk/fred_client/assets"
)

// Source is a single instance of a FRED Source.
type Source struct {
	ID    int    `json:"id" xml:"id,attr"`
	Start string `json:"realtime_start" xml:"realtime_start,attr"`
	End   string `json:"realtime_end" xml:"realtime_end,attr"`
	Name  string `json:"name" xml:"name,attr"`
	Link  string `json:"link" xml:"link,attr"`
}

// GetSources will get all sources of economic data.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/sources.html
func (f *FredClient) GetSources(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, SourcesParam)

	if err != nil {
		f.logError(SourcesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSource will get a source of economic data.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/source.html
func (f *FredClient) GetSource(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, SourceParam)

	if err != nil {
		f.logError(SourceParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSourceReleases will get the releases for a source.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/source_releases.html
func (f *FredClient) GetSourceReleases(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, SourceReleasesParam)

	if err != nil {
		f.logError(SourceReleasesParam, err)
		return nil, err
	}

	return fc, nil

}
