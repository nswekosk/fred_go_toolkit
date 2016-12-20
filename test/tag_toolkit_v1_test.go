package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTags(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		tags, err := xmlFredClient.GetTags(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})
}

func TestGetRelatedTags(t *testing.T) {

	params := make(map[string]interface{})

	params["tag_names"] = "monetary aggregates;weekly"

	Convey("", t, func() {
		tags, err := xmlFredClient.GetRelatedTags(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})
}

func TestGetTagSeries(t *testing.T) {

	params := make(map[string]interface{})

	params["tag_names"] = "slovenia;food;oecd"

	Convey("", t, func() {
		tags, err := xmlFredClient.GetTagSeries(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetTagSeries(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})

}
