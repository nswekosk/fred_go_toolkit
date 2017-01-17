package fred_go_toolkit

// Categories represents a collection of FRED categories.
type Categories struct {
	CategoryCol []Category `json:"categories"`
}

// Category is a single instance of a FRED category.
type Category struct {
	ID       int    `json:"id" xml:"id,attr"`
	Name     string `json:"name" xml:"name,attr"`
	ParentID int    `json:"parent_id" xml:"parent_id,attr"`
}

// GetCategory will get a category.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/category.html
func (f *FredClient) GetCategory(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categoryParam)

	if err != nil {
		f.logError(categoryParam, err)
		return nil, err
	}

	return fc, nil

}

// GetCategoryChildren will get the child categories for a specified parent category.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/category_children.html
func (f *FredClient) GetCategoryChildren(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categoryChildrenParam)

	if err != nil {
		f.logError(categoryChildrenParam, err)
		return nil, err
	}

	return fc, nil

}

// GetRelatedCategory will get the related categories for a category.
// A related category is a one-way relation between 2 categories that is not part of a parent-child category hierarchy.
// Most categories do not have related categories.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/category_related.html
func (f *FredClient) GetRelatedCategory(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categoryRelatedParam)

	if err != nil {
		f.logError(categoryRelatedParam, err)
		return nil, err
	}

	return fc, nil

}

// GetCategorySeries will get the series in a category.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/category_series.html
func (f *FredClient) GetCategorySeries(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categorySeriesParam)

	if err != nil {
		f.logError(categorySeriesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetCategoryTags will get the FRED tags for a category.
// Optionally, filter results by tag name, tag group, or search.
// Series are assigned tags and categories.
// Indirectly through series, it is possible to get the tags for a category.
// No tags exist for a category that does not have series.
// See the related request GetCategoryRelatedTags.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/category_tags.html
func (f *FredClient) GetCategoryTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categoryTagsParam)

	if err != nil {
		f.logError(categoryTagsParam, err)
		return nil, err
	}

	return fc, nil

}

// GetCategoryRelatedTags will get the related FRED tags for one or more FRED tags within a category.
// Optionally, filter results by tag group or search.
//
// FRED tags are attributes assigned to series.
// For this request, related FRED tags are the tags assigned to series that match all tags in the tag_names parameter, no tags in the exclude_tag_names parameter, and the category set by the category_id parameter.
// See the related request GetCategoryTags.
//
// Series are assigned tags and categories.
// Indirectly through series, it is possible to get the tags for a category.
// No tags exist for a category that does not have series.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/category_related_tags.html
func (f *FredClient) GetCategoryRelatedTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, categoryRelatedTagsParam)

	if err != nil {
		f.logError(categoryRelatedTagsParam, err)
		return nil, err
	}

	return fc, nil

}
