package lib

import (
	"encoding/json"
	"errors"
)

type Categories struct {
	CategoryCol []Category `categories`
}

type Category struct {
	id        int    `id`
	name      string `name`
	parent_id string `parent_id`
}

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
