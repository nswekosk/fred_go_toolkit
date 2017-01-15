package lib

import (
	. "github.com/nswekosk/fred_client/assets"
)

// Tag is a single instance of a FRED tag.
type Tag struct {
	Name        string `json:"name" xml:"name,attr"`
	GroupID     string `json:"group_id" xml:"group_id,attr"`
	Notes       string `json:"notes" xml:"notes,attr"`
	Created     string `json:"created" xml:"created,attr"`
	Popularity  int    `json:"popularity" xml:"popularity,attr"`
	SeriesCount int    `json:"series_count" xml:"series_count,attr"`
}

// GetTags will get FRED tags.
// Optionally, filter results by tag name, tag group, or search.
// FRED tags are attributes assigned to series.
// See the related request GetRelatedTags.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/tags.html
func (f *FredClient) GetTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, TagsParam)

	if err != nil {
		f.logError(TagsParam, err)
		return nil, err
	}

	return fc, nil

}

// GetRelatedTags will get the related FRED tags for one or more FRED tags.
// Optionally, filter results by tag group or search.
//
// FRED tags are attributes assigned to series.
// Related FRED tags are the tags assigned to series that match all tags in the tag_names parameter and no tags in the exclude_tag_names parameter.
// See the related request GetTags.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/related_tags.html
func (f *FredClient) GetRelatedTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, RelatedTagsParam)

	if err != nil {
		f.logError(RelatedTagsParam, err)
		return nil, err
	}

	return fc, nil

}

// GetTagSeries will get the series matching all tags in the tag_names parameter and no tags in the exclude_tag_names parameter.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/tags_series.html
func (f *FredClient) GetTagSeries(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, TagsSeriesParam)

	if err != nil {
		f.logError(TagsSeriesParam, err)
		return nil, err
	}

	return fc, nil

}
