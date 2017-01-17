package fred_client

const (
	//////////////////////////////////////////////////
	// fileTypeJSON will return json data from API. //
	//////////////////////////////////////////////////

	FileTypeJSON = "json"

	////////////////////////////////////////////////
	// fileTypeXML will return xml data from API. //
	////////////////////////////////////////////////

	FileTypeXML = "xml"

	//////////////////
	// FRED API URL //
	//////////////////

	apiURL = "https://api.stlouisfed.org/fred"

	////////////
	// errors //
	////////////

	errorNoAPIKey          = "Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account."
	errorNoParameters      = "No parameters were added. Please update you parameter input."
	errorLibraryFail       = "There was an error in processing the query. Please contact the client administrator."
	errorInvalidAPIKey     = "API key is invalid. Please supply a valid API key."
	errorIncorrectFileType = "You must provide a valid file type in order to use this client. Ex: '', 'json', 'xml'"
	////////////////////////////
	// paramLookup attributes //
	////////////////////////////

	paramLookupExt    = "extension"
	paramLookupParams = "params"

	///////////////////////
	// method type param //
	///////////////////////

	categoryParam            = "CATEGORY"
	categoryChildrenParam    = "CATEGORY_CHILDREN"
	categoryRelatedParam     = "CATEGORY_RELATED"
	categorySeriesParam      = "CATEGORY_SERIES"
	categoryTagsParam        = "CATEGORY_TAGS"
	categoryRelatedTagsParam = "CATEGORY_RELATED_TAGS"

	releasesParam           = "RELEASES"
	releasesDatesParam      = "RELEASES_DATES"
	releaseParam            = "RELEASE"
	releaseDatesParam       = "RELEASE_DATES"
	releaseSeriesParam      = "RELEASE_SERIES"
	releaseSourcesParam     = "RELEASE_SOURCES"
	releaseTagsParam        = "RELEASE_TAGS"
	releaseRelatedTagsParam = "RELEASE_RELATED_TAGS"
	releaseTablesParam      = "RELEASE_TABLES"

	seriesParam                  = "SERIES"
	seriesCategoriesParam        = "SERIES_CATEGORIES"
	seriesObservationsParam      = "SERIES_OBSERVATIONS"
	seriesReleaseParam           = "SERIES_RELEASE"
	seriesSearchParam            = "SERIES_SEARCH"
	seriesSearchTagsParam        = "SERIES_SEARCH_TAGS"
	seriesSearchRelatedTagsParam = "SERIES_SEARCH_RELATED_TAGS"
	seriesTagsParam              = "SERIES_TAGS"
	seriesUpdatesParam           = "SERIES_UPDATES"
	seriesVintagedatesParam      = "SERIES_VINTAGEDATES"

	sourcesParam        = "SOURCES"
	sourceParam         = "SOURCE"
	sourceReleasesParam = "SOURCE_RELEASES"

	tagsParam        = "TAGS"
	relatedTagsParam = "RELATED_TAGS"
	tagsSeriesParam  = "TAGS_SERIES"
)
