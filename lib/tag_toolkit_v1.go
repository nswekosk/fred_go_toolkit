package lib

type Tag struct {
	Name        string `json:"name" xml:"name,attr"`
	GroupID     string `json:"group_id" xml:"group_id,attr"`
	Notes       string `json:"notes" xml:"notes,attr"`
	Created     string `json:"created" xml:"created,attr"`
	Popularity  int    `json:"popularity" xml:"popularity,attr"`
	SeriesCount int    `json:"series_count" xml:"series_count,attr"`
}

/********************************
 ** GetTags
 **
 ** Get all tags, search for tags,
 ** or get tags by name.
 ********************************/
func (f *FredClient) GetTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, tags)

	if err != nil {
		f.logError(tags, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetRelatedTags
 **
 ** Get the related tags for one
 ** or more tags.
 ********************************/
func (f *FredClient) GetRelatedTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, relatedTags)

	if err != nil {
		f.logError(relatedTags, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetTagSeries
 **
 ** Get the series matching tags.
 ********************************/
func (f *FredClient) GetTagSeries(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, tagsSeries)

	if err != nil {
		f.logError(tagsSeries, err)
		return nil, err
	}

	return fc, nil

}
