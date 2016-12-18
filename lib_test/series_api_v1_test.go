package lib_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSeries(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srs, err := xmlFredClient.GetSeries(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(srs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetSeries(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(srs, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesCategories(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srsCts, err := xmlFredClient.GetSeriesCategories(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(srsCts, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srsCts, err := jsonFredClient.GetSeriesCategories(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(srsCts, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesObservations(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srsObs, err := xmlFredClient.GetSeriesObservations(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(srsObs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srsObs, err := jsonFredClient.GetSeriesObservations(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(srsObs, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesRelease(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srsRls, err := xmlFredClient.GetSeriesRelease(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(srsRls, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srsRls, err := jsonFredClient.GetSeriesRelease(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(srsRls, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesSearch(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srs, err := xmlFredClient.GetSeriesSearch(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(srs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetSeriesSearch(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(srs, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesSearchTags(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		tags, err := xmlFredClient.GetSeriesSearchTags(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetSeriesSearchTags(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesSearchRelatedTags(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		tags, err := xmlFredClient.GetSeriesSearchRelatedTags(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetSeriesSearchRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesTags(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		tags, err := xmlFredClient.GetSeriesTags(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetSeriesTags(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesUpdates(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srs, err := xmlFredClient.GetSeriesUpdates(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(srs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetSeriesUpdates(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(srs, ShouldNotBeNil)
		})
	})

}

func TestGetSeriesVintageDates(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		vds, err := xmlFredClient.GetSeriesVintageDates(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(vds, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		vds, err := jsonFredClient.GetSeriesVintageDates(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(vds, ShouldNotBeNil)
		})
	})

}
