package lib

import (
	"encoding/json"
	"errors"
)

type Tags struct {
	Start     string   `realtime_start`
	End       string   `realtime_end`
	OrderBy   string   `order_by`
	SortOrder string   `sort_order`
	Count     int      `count`
	Offset    int      `offset`
	Limit     int      `limit`
	Tags      []Tag    `tag`
	Sources   []Source `sources`
}

type Tag struct {
	Name        string `name`
	GroupId     string `group_id`
	Notes       string `notes`
	Created     string `created`
	Popularity  int    `popularity`
	SeriesCount int    `series_count`
}

/********************************
 ** GetTags
 **
 ** Get all tags, search for tags,
 ** or get tags by name.
 ********************************/
func (f *FredClient) GetTags(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "TAGS")

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
 ** GetRelatedTags
 **
 ** Get the related tags for one
 ** or more tags.
 ********************************/
func (f *FredClient) GetRelatedTags(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "RELATED_TAGS")

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
 ** GetTagSeries
 **
 ** Get the series matching tags.
 ********************************/
func (f *FredClient) GetTagSeries(params map[string]interface{}) (*Seriess, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	srs := &Seriess{}

	resp, err := f.callAPI(params, "RELATED_TAGS")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(srs)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return srs, nil

}
