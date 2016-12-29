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
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 100)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 100)

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
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 13)
			So(tags.Tags[0].GroupID, ShouldContainSubstring, "geot")
			So(tags.Tags[0].Name, ShouldContainSubstring, "nation")

		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 13)
			So(tags.Tags[0].GroupID, ShouldContainSubstring, "geot")
			So(tags.Tags[0].Name, ShouldContainSubstring, "nation")
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
			So(len(tags.Seriess), ShouldBeGreaterThanOrEqualTo, 18)
			So(tags.Seriess[0].ID, ShouldContainSubstring, "CPGDFD02SIA657N")
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetTagSeries(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Seriess), ShouldBeGreaterThanOrEqualTo, 18)
			So(tags.Seriess[0].ID, ShouldContainSubstring, "CPGDFD02SIA657N")
		})
	})

}
