package test

import (
	"os"
	"testing"

	. "github.com/nswekosk/fred_client/lib"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	apiKey         = ""
	xmlFredClient  = &FredClient{}
	jsonFredClient = &FredClient{}
)

func TestMain(m *testing.M) {

	apiKey = os.Getenv("FRED_API_KEY")

	var err error

	xmlFredClient, err = CreateFredClient(apiKey, FileTypeXML)

	if err != nil {
		panic(err.Error())
	}

	jsonFredClient, err = CreateFredClient(apiKey, FileTypeJSON)
	if err != nil {
		panic(err.Error())
	}

	ret := m.Run()
	os.Exit(ret)
}

func TestUpdateAPIKEY(t *testing.T) {

	key := ""

	Convey("", t, func() {
		err := xmlFredClient.UpdateAPIKEY(key)

		So(err, ShouldNotBeNil)
	})

	key = "23235234215423452345324"

	Convey("", t, func() {
		err := jsonFredClient.UpdateAPIKEY(key)

		So(err, ShouldNotBeNil)
	})

}
