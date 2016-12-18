package lib_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSources(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srcs, err := xmlFredClient.GetSources(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(srcs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srcs, err := jsonFredClient.GetSources(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(srcs, ShouldNotBeNil)
		})
	})

}

func TestGetSource(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		src, err := xmlFredClient.GetSource(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(src, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		src, err := jsonFredClient.GetSource(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(src, ShouldNotBeNil)
		})
	})

}

func TestGetSourceReleases(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		rls, err := xmlFredClient.GetSourceReleases(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(rls, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		rls, err := jsonFredClient.GetSourceReleases(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(rls, ShouldNotBeNil)
		})
	})

}
