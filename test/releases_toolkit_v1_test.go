package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetReleases(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		rls, err := xmlFredClient.GetReleases(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(rls, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		rls, err := jsonFredClient.GetReleases(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rls, ShouldNotBeNil)
		})
	})

}

func TestGetReleasesDates(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		dts, err := xmlFredClient.GetReleasesDates(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(dts, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		dts, err := jsonFredClient.GetReleasesDates(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(dts, ShouldNotBeNil)
		})
	})

}

func TestGetRelease(t *testing.T) {

	params := make(map[string]interface{})

	params["release_id"] = 53

	Convey("", t, func() {
		rls, err := xmlFredClient.GetRelease(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(rls, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		rls, err := jsonFredClient.GetRelease(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rls, ShouldNotBeNil)
		})
	})

}

func TestGetReleaseDates(t *testing.T) {

	params := make(map[string]interface{})

	params["release_id"] = 82

	Convey("", t, func() {
		rlsDts, err := xmlFredClient.GetReleaseDates(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(rlsDts, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		rlsDts, err := jsonFredClient.GetReleaseDates(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rlsDts, ShouldNotBeNil)
		})
	})

}

func TestGetReleaseSeries(t *testing.T) {

	params := make(map[string]interface{})

	params["release_id"] = 51

	Convey("", t, func() {
		srs, err := xmlFredClient.GetReleaseSeries(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(srs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetReleaseSeries(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
		})
	})

}

func TestGetReleaseSources(t *testing.T) {

	params := make(map[string]interface{})

	params["release_id"] = 51

	Convey("", t, func() {
		srs, err := xmlFredClient.GetReleaseSources(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(srs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetReleaseSources(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
		})
	})

}

func TestGetReleaseTags(t *testing.T) {

	params := make(map[string]interface{})

	params["release_id"] = 86

	Convey("", t, func() {
		tags, err := xmlFredClient.GetReleaseTags(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetReleaseTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})

}

func TestGetReleaseRelatedTags(t *testing.T) {

	params := make(map[string]interface{})

	params["release_id"] = 86
	params["tag_names"] = "sa;foreign"

	Convey("", t, func() {
		tags, err := xmlFredClient.GetReleaseRelatedTags(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetReleaseRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
		})
	})

}
