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
			So(len(rls.Releases), ShouldBeGreaterThanOrEqualTo, 158)
			So(rls.Releases[0].Name, ShouldContainSubstring, "Advance Monthly Sales for Retail and Food Services")

		})
	})

	Convey("", t, func() {
		rls, err := jsonFredClient.GetReleases(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rls, ShouldNotBeNil)
			So(len(rls.Releases), ShouldBeGreaterThanOrEqualTo, 158)
			So(rls.Releases[0].Name, ShouldContainSubstring, "Advance Monthly Sales for Retail and Food Services")
		})
	})

}

func TestGetReleasesDates(t *testing.T) {

	params := make(map[string]interface{})

	params["release_id"] = 82

	Convey("", t, func() {
		dts, err := xmlFredClient.GetReleasesDates(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(dts, ShouldNotBeNil)
			So(dts.ReleaseDates[0].Date, ShouldContainSubstring, "1997-02-10")
		})
	})

	Convey("", t, func() {
		dts, err := jsonFredClient.GetReleasesDates(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(dts, ShouldNotBeNil)
			So(dts.ReleaseDates[0].Date, ShouldContainSubstring, "1997-02-10")
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
			So(rls.Releases[0].ID, ShouldBeGreaterThanOrEqualTo, 53)
			So(rls.Releases[0].Name, ShouldContainSubstring, "Gross Domestic Product")
		})
	})

	Convey("", t, func() {
		rls, err := jsonFredClient.GetRelease(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rls, ShouldNotBeNil)
			So(len(rls.Releases), ShouldBeGreaterThanOrEqualTo, 1)
			So(rls.Releases[0].ID, ShouldBeGreaterThanOrEqualTo, 53)
			So(rls.Releases[0].Name, ShouldContainSubstring, "Gross Domestic Product")
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
			So(len(rlsDts.ReleaseDates), ShouldBeGreaterThanOrEqualTo, 17)

		})
	})

	Convey("", t, func() {
		rlsDts, err := jsonFredClient.GetReleaseDates(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rlsDts, ShouldNotBeNil)
			So(len(rlsDts.ReleaseDates), ShouldBeGreaterThanOrEqualTo, 17)

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
			So(len(srs.Seriess), ShouldBeGreaterThanOrEqualTo, 39)
			So(srs.Seriess[0].ID, ShouldBeGreaterThanOrEqualTo, "BOMTVLM133S")
			So(srs.Seriess[0].Title, ShouldBeGreaterThanOrEqualTo, "U.S. Imports of Services - Travel")

		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetReleaseSeries(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
			So(len(srs.Seriess), ShouldBeGreaterThanOrEqualTo, 39)
			So(srs.Seriess[0].ID, ShouldBeGreaterThanOrEqualTo, "BOMTVLM133S")
			So(srs.Seriess[0].Title, ShouldBeGreaterThanOrEqualTo, "U.S. Imports of Services - Travel")
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
			So(len(srs.Sources), ShouldBeGreaterThanOrEqualTo, 2)
			So(srs.Sources[0].ID, ShouldBeGreaterThanOrEqualTo, 18)
			So(srs.Sources[0].Name, ShouldContainSubstring, "U.S. Bureau of Economic Analysis")

		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetReleaseSources(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
			So(len(srs.Sources), ShouldBeGreaterThanOrEqualTo, 2)
			So(srs.Sources[0].ID, ShouldBeGreaterThanOrEqualTo, 18)
			So(srs.Sources[0].Name, ShouldContainSubstring, "U.S. Bureau of Economic Analysis")
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
			So(tags.Tags[0].GroupID, ShouldContainSubstring, "gen")
			So(tags.Tags[0].Name, ShouldContainSubstring, "commercial")
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetReleaseTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(tags.Tags[0].GroupID, ShouldContainSubstring, "gen")
			So(tags.Tags[0].Name, ShouldContainSubstring, "commercial")
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
			So(tags.Tags[0].GroupID, ShouldContainSubstring, "gen")
			So(tags.Tags[0].Name, ShouldContainSubstring, "commercial")
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetReleaseRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(tags.Tags[0].GroupID, ShouldContainSubstring, "gen")
			So(tags.Tags[0].Name, ShouldContainSubstring, "commercial")
		})
	})

}

func TestGetReleaseTables(t *testing.T) {

	params := make(map[string]interface{})

	params["release_id"] = 53
	params["element_id"] = 12886

	Convey("", t, func() {
		res, err := xmlFredClient.GetReleaseTables(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(res, ShouldNotBeBlank)
		})
	})

	Convey("", t, func() {
		res, err := jsonFredClient.GetReleaseTables(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(res, ShouldNotBeBlank)
		})
	})

}
