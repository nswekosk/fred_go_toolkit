package fred_go_toolkit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSeries(t *testing.T) {

	params := make(map[string]interface{})

	params["series_id"] = "GNPCA"

	Convey("", t, func() {
		srs, err := xmlFredClient.GetSeries(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(srs, ShouldNotBeNil)
			So(srs.Seriess[0].ID, ShouldBeGreaterThanOrEqualTo, "GNPCA")
			So(srs.Seriess[0].Title, ShouldBeGreaterThanOrEqualTo, "Real Gross National Product")

		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetSeries(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
			So(srs.Seriess[0].ID, ShouldBeGreaterThanOrEqualTo, "GNPCA")
			So(srs.Seriess[0].Title, ShouldBeGreaterThanOrEqualTo, "Real Gross National Product")
		})
	})

}

func TestGetSeriesCategories(t *testing.T) {

	params := make(map[string]interface{})

	params["series_id"] = "EXJPUS"

	Convey("", t, func() {
		srsCts, err := xmlFredClient.GetSeriesCategories(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(srsCts, ShouldNotBeNil)
			So(len(srsCts.Categories), ShouldBeGreaterThanOrEqualTo, 2)
			So(srsCts.Categories[0].ID, ShouldBeGreaterThanOrEqualTo, 95)
			So(srsCts.Categories[0].Name, ShouldContainSubstring, "Monthly Rates")
		})
	})

	Convey("", t, func() {
		srsCts, err := jsonFredClient.GetSeriesCategories(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srsCts, ShouldNotBeNil)
			So(len(srsCts.Categories), ShouldBeGreaterThanOrEqualTo, 2)
			So(srsCts.Categories[0].ID, ShouldBeGreaterThanOrEqualTo, 95)
			So(srsCts.Categories[0].Name, ShouldContainSubstring, "Monthly Rates")
		})
	})

}

func TestGetSeriesObservations(t *testing.T) {

	params := make(map[string]interface{})

	params["series_id"] = "GNPCA"

	Convey("", t, func() {
		srsObs, err := xmlFredClient.GetSeriesObservations(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(srsObs, ShouldNotBeNil)
			So(len(srsObs.Observations), ShouldBeGreaterThanOrEqualTo, 84)

		})
	})

	Convey("", t, func() {
		srsObs, err := jsonFredClient.GetSeriesObservations(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srsObs, ShouldNotBeNil)
			So(len(srsObs.Observations), ShouldBeGreaterThanOrEqualTo, 84)

		})
	})

}

func TestGetSeriesRelease(t *testing.T) {

	params := make(map[string]interface{})

	params["series_id"] = "IRA"

	Convey("", t, func() {
		srsRls, err := xmlFredClient.GetSeriesRelease(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(srsRls, ShouldNotBeNil)
			So(len(srsRls.Releases), ShouldBeGreaterThanOrEqualTo, 1)
			So(srsRls.Releases[0].Name, ShouldContainSubstring, "H.6 Money Stock Measures")
			So(srsRls.Releases[0].ID, ShouldBeGreaterThanOrEqualTo, 21)

		})
	})

	Convey("", t, func() {
		srsRls, err := jsonFredClient.GetSeriesRelease(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srsRls, ShouldNotBeNil)
			So(len(srsRls.Releases), ShouldBeGreaterThanOrEqualTo, 1)
			So(srsRls.Releases[0].Name, ShouldContainSubstring, "H.6 Money Stock Measures")
			So(srsRls.Releases[0].ID, ShouldBeGreaterThanOrEqualTo, 21)
		})
	})

}

func TestGetSeriesSearch(t *testing.T) {

	params := make(map[string]interface{})

	params["search_text"] = "monetary service index"

	Convey("", t, func() {
		srs, err := xmlFredClient.GetSeriesSearch(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(srs, ShouldNotBeNil)
			So(len(srs.Seriess), ShouldBeGreaterThanOrEqualTo, 25)
			So(srs.Seriess[0].ID, ShouldContainSubstring, "MSIM2")
			So(srs.Seriess[0].Title, ShouldContainSubstring, "Monetary Services Index: M2 (preferred)")
		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetSeriesSearch(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
			So(len(srs.Seriess), ShouldBeGreaterThanOrEqualTo, 25)
			So(srs.Seriess[0].ID, ShouldContainSubstring, "MSIM2")
			So(srs.Seriess[0].Title, ShouldContainSubstring, "Monetary Services Index: M2 (preferred)")
		})
	})

}

func TestGetSeriesSearchTags(t *testing.T) {

	params := make(map[string]interface{})

	params["series_search_text"] = "monetary service index"

	Convey("", t, func() {
		tags, err := xmlFredClient.GetSeriesSearchTags(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 18)

		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetSeriesSearchTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 18)

		})
	})

}

func TestGetSeriesSearchRelatedTags(t *testing.T) {

	params := make(map[string]interface{})

	params["series_search_text"] = "mortgage rate"
	params["tag_names"] = "30-year;frb"

	Convey("", t, func() {
		tags, err := xmlFredClient.GetSeriesSearchRelatedTags(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 10)
			So(tags.Tags[0].Name, ShouldContainSubstring, "conventional")
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetSeriesSearchRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 10)
			So(tags.Tags[0].Name, ShouldContainSubstring, "conventional")
		})
	})

}

func TestGetSeriesTags(t *testing.T) {

	params := make(map[string]interface{})

	params["series_id"] = "STLFSI"

	Convey("", t, func() {
		tags, err := xmlFredClient.GetSeriesTags(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 8)
			So(tags.Tags[0].Name, ShouldContainSubstring, "nsa")

		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetSeriesTags(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(tags, ShouldNotBeNil)
			So(len(tags.Tags), ShouldBeGreaterThanOrEqualTo, 8)
			So(tags.Tags[0].Name, ShouldContainSubstring, "nsa")

		})
	})

}

func TestGetSeriesUpdates(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srs, err := xmlFredClient.GetSeriesUpdates(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(srs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetSeriesUpdates(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srs, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesVintageDates(t *testing.T) {

	params := make(map[string]interface{})

	params["series_id"] = "GNPCA"

	Convey("", t, func() {
		vds, err := xmlFredClient.GetSeriesVintageDates(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(vds, ShouldNotBeNil)
			So(len(vds.VintageDates), ShouldBeGreaterThanOrEqualTo, 162)

		})
	})

	Convey("", t, func() {
		vds, err := jsonFredClient.GetSeriesVintageDates(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(vds, ShouldNotBeNil)
			So(len(vds.VintageDates), ShouldBeGreaterThanOrEqualTo, 162)

		})
	})

}
