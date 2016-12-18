package lib_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTags(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		tags, err := xmlFredClient.GetTags(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetTags(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})
}

func TestGetRelatedTags(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		tags, err := xmlFredClient.GetRelatedTags(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})
}

func TestGetTagSeries(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		tags, err := xmlFredClient.GetTagSeries(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetTagSeries(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

}
