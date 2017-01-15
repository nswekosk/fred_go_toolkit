package lib

const (
	//////////////////////////////////////////////////
	// FileTypeJSON will return json data from API. //
	//////////////////////////////////////////////////

	FileTypeJSON = "json"

	////////////////////////////////////////////////
	// FileTypeXML will return xml data from API. //
	////////////////////////////////////////////////

	FileTypeXML = "xml"

	//////////////////
	// FRED API URL //
	//////////////////

	ApiURL = "https://api.stlouisfed.org/fred"

	////////////
	// errors //
	////////////

	ErrorNoAPIKey          = "Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account."
	ErrorNoParameters      = "No parameters were added. Please update you parameter input."
	ErrorLibraryFail       = "There was an error in processing the query. Please contact the client administrator."
	ErrorInvalidAPIKey     = "API key is invalid. Please supply a valid API key."
	ErrorIncorrectFileType = "You must provide a valid file type in order to use this client. Ex: '', 'json', 'xml'"
	////////////////////////////
	// paramLookup attributes //
	////////////////////////////

	ParamLookupExt    = "extension"
	ParamLookupParams = "params"

	///////////////////////
	// method type param //
	///////////////////////

	CategoryParam            = "CATEGORY"
	CategoryChildrenParam    = "CATEGORY_CHILDREN"
	CategoryRelatedParam     = "CATEGORY_RELATED"
	CategorySeriesParam      = "CATEGORY_SERIES"
	CategoryTagsParam        = "CATEGORY_TAGS"
	CategoryRelatedTagsParam = "CATEGORY_RELATED_TAGS"

	ReleasesParam           = "RELEASES"
	ReleasesDatesParam      = "RELEASES_DATES"
	ReleaseParam            = "RELEASE"
	ReleaseDatesParam       = "RELEASE_DATES"
	ReleaseSeriesParam      = "RELEASE_SERIES"
	ReleaseSourcesParam     = "RELEASE_SOURCES"
	ReleaseTagsParam        = "RELEASE_TAGS"
	ReleaseRelatedTagsParam = "RELEASE_RELATED_TAGS"
	ReleaseTablesParam      = "RELEASE_TABLES"

	SeriesParam                  = "SERIES"
	SeriesCategoriesParam        = "SERIES_CATEGORIES"
	SeriesObservationsParam      = "SERIES_OBSERVATIONS"
	SeriesReleaseParam           = "SERIES_RELEASE"
	SeriesSearchParam            = "SERIES_SEARCH"
	SeriesSearchTagsParam        = "SERIES_SEARCH_TAGS"
	SeriesSearchRelatedTagsParam = "SERIES_SEARCH_RELATED_TAGS"
	SeriesTagsParam              = "SERIES_TAGS"
	SeriesUpdatesParam           = "SERIES_UPDATES"
	SeriesVintagedatesParam      = "SERIES_VINTAGEDATES"

	SourcesParam        = "SOURCES"
	SourceParam         = "SOURCE"
	SourceReleasesParam = "SOURCE_RELEASES"

	TagsParam        = "TAGS"
	RelatedTagsParam = "RELATED_TAGS"
	TagsSeriesParam  = "TAGS_SERIES"
)
