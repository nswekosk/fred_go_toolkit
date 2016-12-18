package lib

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Category struct {
	id        int    `id`
	name      string `name`
	parent_id string `parent_id`
}

func (f *FredClient) GetCategory(CategoryID string) (*Category, error) {

	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	if CategoryID == "" {
		return nil, errors.New("You must provide a cateory id to retrieve a category. CategoryID's can be found at ")
	}

	cat := &Category{}
	url := f.requestUrl + "/category_id?=" + CategoryID

	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.New("There was an error in processing the query. Please contact the client administrator.")
	}

	err = json.NewDecoder(resp.Body).Decode(cat)

	if err != nil {
		return nil, errors.New("There was an error in processing the query. Please contact the client administrator.")
	}

	return cat, nil

}

func (f *FredClient) GetCategoryChildren(params ...string) (*[]Category, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	return nil, nil

}

func (f *FredClient) GetRelatedCategory(params ...string) (*[]Category, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	return nil, nil

}

func (f *FredClient) GetCategorySeries(params ...interface{}) (*Seriess, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	return nil, nil

}

func (f *FredClient) GetCategoryTags(params ...interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	return nil, nil

}

func (f *FredClient) GetCategoryRelatedTags(params ...interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	return nil, nil

}
