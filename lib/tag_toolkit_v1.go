package lib

type Tags struct {
	Start     string   `json:"realtime_start" xml:"realtime_start"`
	End       string   `json:"realtime_end" xml:"realtime_end"`
	OrderBy   string   `json:"order_by" xml:"order_by"`
	SortOrder string   `json:"sort_order" xml:"sort_order"`
	Count     int      `json:"count" xml:"count"`
	Offset    int      `json:"offset" xml:"offset"`
	Limit     int      `json:"limit" xml:"limit"`
	Tags      []Tag    `json:"tag" xml:"tag"`
	Sources   []Source `json:"sources" xml:"sources"`
}

type Tag struct {
	Name        string `json:"name" xml:"name"`
	GroupID     string `json:"group_id" xml:"group_id"`
	Notes       string `json:"notes" xml:"notes"`
	Created     string `json:"created" xml:"created"`
	Popularity  int    `json:"popularity" xml:"popularity"`
	SeriesCount int    `json:"series_count" xml:"series_count"`
}

/********************************
 ** GetTags
 **
 ** Get all tags, search for tags,
 ** or get tags by name.
 ********************************/
func (f *FredClient) GetTags(params map[string]interface{}) (*Tags, error) {

	tags, err := f.operate(params, tags)

	if err != nil {
		return nil, err
	}

	return (tags.(*Tags)), nil

}

/********************************
 ** GetRelatedTags
 **
 ** Get the related tags for one
 ** or more tags.
 ********************************/
func (f *FredClient) GetRelatedTags(params map[string]interface{}) (*Tags, error) {

	tags, err := f.operate(params, relatedTags)

	if err != nil {
		return nil, err
	}

	return (tags.(*Tags)), nil

}

/********************************
 ** GetTagSeries
 **
 ** Get the series matching tags.
 ********************************/
func (f *FredClient) GetTagSeries(params map[string]interface{}) (*Seriess, error) {

	srs, err := f.operate(params, tagsSeries)

	if err != nil {
		return nil, err
	}

	return (srs.(*Seriess)), nil

}
