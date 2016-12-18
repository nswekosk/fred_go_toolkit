package lib

type Tags struct {
	Start     string   `json:"realtime_start"`
	End       string   `json:"realtime_end"`
	OrderBy   string   `json:"order_by"`
	SortOrder string   `json:"sort_order"`
	Count     int      `json:"count"`
	Offset    int      `json:"offset"`
	Limit     int      `json:"limit"`
	Tags      []Tag    `json:"tag"`
	Sources   []Source `json:"sources"`
}

type Tag struct {
	Name        string `json:"name"`
	GroupID     string `json:"group_id"`
	Notes       string `json:"notes"`
	Created     string `json:"created"`
	Popularity  int    `json:"popularity"`
	SeriesCount int    `json:"series_count"`
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
