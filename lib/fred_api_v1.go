package lib

type FredClient struct {
	API_KEY    string
	Categories Category
	Releases   Releases
	Series     Series
	Sources    Seriess
	Tags       Tags
}

func CreateClient(ApiKey string) FredClient {

	return FredClient{
		API_KEY: ApiKey,
	}

}
