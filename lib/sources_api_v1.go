package lib

type Sources struct {
	Start      string   `json:"realtime_start"`
	End        string   `json:"realtime_end"`
	OrderBy    string   `json:"order_by"`
	SortOrder  string   `json:"sort_order"`
	Count      int      `json:"count"`
	Offset     int      `json:"offset"`
	Limit      int      `json:"limit"`
	SourcesCol []Source `json:"sources"`
}

type Source struct {
	ID    int    `json:"id"`
	Start string `json:"realtime_start"`
	End   string `json:"realtime_end"`
	Name  string `json:"name"`
	Link  string `json:"link"`
}

/********************************
 ** GetSources
 **
 ** Get all sources of economic
 ** data.
 ********************************/
func (f *FredClient) GetSources(params map[string]interface{}) (*Sources, error) {

	srcs, err := f.operate(params, sources)

	if err != nil {
		return nil, err
	}

	return (srcs.(*Sources)), nil

}

/********************************
 ** GetSource
 **
 ** Get a source of economic data.
 ********************************/
func (f *FredClient) GetSource(params map[string]interface{}) (*Sources, error) {

	srcs, err := f.operate(params, source)

	if err != nil {
		return nil, err
	}

	return (srcs.(*Sources)), nil

}

/********************************
 ** GetSourceReleases
 **
 ** Get the releases for a source.
 ********************************/
func (f *FredClient) GetSourceReleases(params map[string]interface{}) (*Releases, error) {

	rls, err := f.operate(params, sourceReleases)

	if err != nil {
		return nil, err
	}

	return (rls.(*Releases)), nil

}
