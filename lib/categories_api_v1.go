package lib

type Categories struct {
	CategoryCol []Category `json:"categories"`
}

type Category struct {
	id        int    `json:"id"`
	name      string `json:"name"`
	parent_id string `json:"parent_id"`
}

/********************************
 ** GetCategory
 **
 ** Get a category.
 ********************************/
func (f *FredClient) GetCategory(params map[string]interface{}) (*Categories, error) {

	cats, err := f.operate(params, "CATEGORY")

	if err != nil {
		return nil, err
	}

	return (cats.(*Categories)), nil

}

/********************************
 ** GetCategoryChildren
 **
 ** Get the child categories for
 ** a specified parent category
 ********************************/
func (f *FredClient) GetCategoryChildren(params map[string]interface{}) (*Categories, error) {

	cats, err := f.operate(params, "CATEGORY_CHILDREN")

	if err != nil {
		return nil, err
	}

	return (cats.(*Categories)), nil

}

/********************************
 ** GetRelatedCategory
 **
 ** Get the related categories
 ** for a category.
 ********************************/
func (f *FredClient) GetRelatedCategory(params map[string]interface{}) (*Categories, error) {

	cats, err := f.operate(params, "CATEGORY_RELATED")

	if err != nil {
		return nil, err
	}

	return (cats.(*Categories)), nil

}

/********************************
 ** GetCategorySeries
 **
 ** Get the series in a category.
 ********************************/
func (f *FredClient) GetCategorySeries(params map[string]interface{}) (*Seriess, error) {

	sers, err := f.operate(params, "CATEGORY_SERIES")

	if err != nil {
		return nil, err
	}

	return (sers.(*Seriess)), nil

}

/********************************
 ** GetCategoryTags
 **
 ** Get the tags for a category.
 ********************************/
func (f *FredClient) GetCategoryTags(params map[string]interface{}) (*Tags, error) {

	tags, err := f.operate(params, "CATEGORY_TAGS")

	if err != nil {
		return nil, err
	}

	return (tags.(*Tags)), nil

}

/********************************
 ** GetCategoryRelatedTags
 **
 ** Get the related tags for a
 ** category.
 ********************************/
func (f *FredClient) GetCategoryRelatedTags(params map[string]interface{}) (*Tags, error) {

	tags, err := f.operate(params, "CATEGORY_RELATED_TAGS")

	if err != nil {
		return nil, err
	}

	return (tags.(*Tags)), nil

}
