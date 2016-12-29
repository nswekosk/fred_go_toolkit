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
			//So(len(rls.Releases), ShouldBeGreaterThanOrEqualTo, 40)
		})
	})

	Convey("", t, func() {
		rls, err := jsonFredClient.GetReleases(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rls, ShouldNotBeNil)
			//So(len(rls.Releases), ShouldBeGreaterThanOrEqualTo, 1)

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
			//So(len(dts.ReleaseDates), ShouldBeGreaterThanOrEqualTo, 40)

		})
	})

	Convey("", t, func() {
		dts, err := jsonFredClient.GetReleasesDates(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(dts, ShouldNotBeNil)
			//So(len(dts.ReleaseDates), ShouldBeGreaterThanOrEqualTo, 40)

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
			So(len(rls.Releases), ShouldBeGreaterThanOrEqualTo, 1)

		})
	})

	Convey("", t, func() {
		rls, err := jsonFredClient.GetRelease(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rls, ShouldNotBeNil)
			So(len(rls.Releases), ShouldBeGreaterThanOrEqualTo, 1)

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
			So(len(rlsDts.ReleaseDates), ShouldBeGreaterThanOrEqualTo, 1)

		})
	})

	Convey("", t, func() {
		rlsDts, err := jsonFredClient.GetReleaseDates(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rlsDts, ShouldNotBeNil)
			So(len(rlsDts.ReleaseDates), ShouldBeGreaterThanOrEqualTo, 1)

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
			So(len(srs.Seriess), ShouldBeGreaterThanOrEqualTo, 1)

		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetReleaseSeries(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
			So(len(srs.Seriess), ShouldBeGreaterThanOrEqualTo, 1)

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
			So(len(srs.Sources), ShouldBeGreaterThanOrEqualTo, 1)

		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetReleaseSources(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
			So(len(srs.Sources), ShouldBeGreaterThanOrEqualTo, 1)

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
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 1)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetReleaseTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 1)
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
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 1)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetReleaseRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 1)
		})
	})

}
