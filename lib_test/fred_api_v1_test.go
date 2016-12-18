package lib_test

import (
	"os"
	"testing"

	. "github.com/Sweko/fred_client/lib"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	apiKey         = ""
	xmlFredClient  = &FredClient{}
	jsonFredClient = &FredClient{}
)

func main() {
	apiKey = os.Getenv("FRED_API_KEY")

	var err error

	xmlFredClient, err = CreateClient(apiKey, FileTypeXML)

	if err != nil {
		panic(err)
	}

	jsonFredClient, err = CreateClient(apiKey, FileTypeJSON)
	if err != nil {
		panic(err)
	}
}

func TestUpdateAPIKEY(t *testing.T) {

	key := ""

	Convey("", t, func() {
		xmlFredClient.UpdateAPIKEY(key)
	})

	Convey("", t, func() {
		jsonFredClient.UpdateAPIKEY(key)
	})

}
