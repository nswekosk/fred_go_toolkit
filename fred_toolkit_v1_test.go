package fred_client

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	apiKey         = ""
	xmlFredClient  = &FredClient{}
	jsonFredClient = &FredClient{}
)

func TestMain(m *testing.M) {

	apiKey = os.Getenv("FRED_API_KEY")

	fredConfig := FredConfig{
		APIKey:   apiKey,
		FileType: FileTypeXML,
		LogFile:  "",
	}

	var err error

	xmlFredClient, err = CreateFredClient(fredConfig)

	if err != nil {
		panic(err.Error())
	}

	fredConfig.FileType = FileTypeJSON

	jsonFredClient, err = CreateFredClient(fredConfig)
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
