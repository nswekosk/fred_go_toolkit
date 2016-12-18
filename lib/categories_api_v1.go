package lib

import (
	"encoding/json"
	"errors"
)

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

	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	cat := &Categories{}

	resp, err := f.callAPI(params, "CATEGORY")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(cat)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return cat, nil

}

/********************************
 ** GetCategoryChildren
 **
 ** Get the child categories for
 ** a specified parent category
 ********************************/
func (f *FredClient) GetCategoryChildren(params map[string]interface{}) (*Categories, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	cats := &Categories{}

	resp, err := f.callAPI(params, "CATEGORY_CHILDREN")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(cats)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return cats, nil

}

/********************************
 ** GetRelatedCategory
 **
 ** Get the related categories
 ** for a category.
 ********************************/
func (f *FredClient) GetRelatedCategory(params map[string]interface{}) (*Categories, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	cats := &Categories{}

	resp, err := f.callAPI(params, "CATEGORY_RELATED")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(cats)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return cats, nil

}

/********************************
 ** GetCategorySeries
 **
 ** Get the series in a category.
 ********************************/
func (f *FredClient) GetCategorySeries(params map[string]interface{}) (*Seriess, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	sers := &Seriess{}

	resp, err := f.callAPI(params, "CATEGORY_SERIES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(sers)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return sers, nil

}

/********************************
 ** GetCategoryTags
 **
 ** Get the tags for a category.
 ********************************/
func (f *FredClient) GetCategoryTags(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "CATEGORY_TAGS")

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
 ** GetCategoryRelatedTags
 **
 ** Get the related tags for a
 ** category.
 ********************************/
func (f *FredClient) GetCategoryRelatedTags(params map[string]interface{}) (*Tags, error) {

	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "CATEGORY_RELATED_TAGS")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(tags)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return tags, nil

}
