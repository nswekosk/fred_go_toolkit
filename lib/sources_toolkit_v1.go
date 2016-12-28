package lib

type Source struct {
	ID    int    `json:"id" xml:"id,attr"`
	Start string `json:"realtime_start" xml:"realtime_start,attr"`
	End   string `json:"realtime_end" xml:"realtime_end,attr"`
	Name  string `json:"name" xml:"name,attr"`
	Link  string `json:"link" xml:"link,attr"`
}

/********************************
 ** GetSources
 **
 ** Get all sources of economic
 ** data.
 ********************************/
func (f *FredClient) GetSources(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, sources)

	if err != nil {
		f.logError(sources, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSource
 **
 ** Get a source of economic data.
 ********************************/
func (f *FredClient) GetSource(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, source)

	if err != nil {
		f.logError(source, err)
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetSourceReleases
 **
 ** Get the releases for a source.
 ********************************/
func (f *FredClient) GetSourceReleases(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, sourceReleases)

	if err != nil {
		f.logError(sourceReleases, err)
		return nil, err
	}

	return fc, nil

}
