package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSources(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srcs, err := xmlFredClient.GetSources(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(srcs, ShouldNotBeNil)
			So(len(srcs.Sources), ShouldBeGreaterThanOrEqualTo, 58)
			So(srcs.Sources[0].Name, ShouldContainSubstring, "Board of Governors of the Federal Reserve System")
		})
	})

	Convey("", t, func() {
		srcs, err := jsonFredClient.GetSources(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(srcs, ShouldNotBeNil)
			So(len(srcs.Sources), ShouldBeGreaterThanOrEqualTo, 58)
			So(srcs.Sources[0].Name, ShouldContainSubstring, "Board of Governors of the Federal Reserve System")
		})
	})

}

func TestGetSource(t *testing.T) {

	params := make(map[string]interface{})

	params["source_id"] = 1

	Convey("", t, func() {
		src, err := xmlFredClient.GetSource(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(src, ShouldNotBeNil)
			So(len(src.Sources), ShouldBeGreaterThanOrEqualTo, 1)
			So(src.Sources[0].Name, ShouldContainSubstring, "Board of Governors of the Federal Reserve System")

		})
	})

	Convey("", t, func() {
		src, err := jsonFredClient.GetSource(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(src, ShouldNotBeNil)
			So(len(src.Sources), ShouldBeGreaterThanOrEqualTo, 1)
			So(src.Sources[0].Name, ShouldContainSubstring, "Board of Governors of the Federal Reserve System")

		})
	})

}

func TestGetSourceReleases(t *testing.T) {

	params := make(map[string]interface{})

	params["source_id"] = 1

	Convey("", t, func() {
		rls, err := xmlFredClient.GetSourceReleases(params)
		So(err, ShouldBeNil)

		Convey("", func() {
			So(rls, ShouldNotBeNil)
			So(len(rls.Releases), ShouldBeGreaterThanOrEqualTo, 26)
			So(rls.Releases[0].Name, ShouldContainSubstring, "G.17 Industrial Production and Capacity Utilization")

		})
	})

	Convey("", t, func() {
		rls, err := jsonFredClient.GetSourceReleases(params)

		So(err, ShouldBeNil)
		Convey("", func() {
			So(rls, ShouldNotBeNil)
			So(len(rls.Releases), ShouldBeGreaterThanOrEqualTo, 26)
			So(rls.Releases[0].Name, ShouldContainSubstring, "G.17 Industrial Production and Capacity Utilization")

		})
	})

}
