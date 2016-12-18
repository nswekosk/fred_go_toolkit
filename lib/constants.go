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

	apiURL = "https://api.stlouisfed.org/fred"

	////////////
	// errors //
	////////////

	errorNoParams     = "Operation may not be performed witout parameters. Refer to 'https://research.stlouisfed.org/docs/api/fred/' for method-specific parameters."
	errorNoAPIKey     = "Operation may not be performed without an APIKEY. APIKEY's can be retrieved at your research.stlouisfed.org user account."
	errorNoParameters = "No parameters were added. Please update you parameter input."
	errorLibraryFail  = "There was an error in processing the query. Please contact the client administrator."

	////////////////////////////
	// paramLookup attributes //
	////////////////////////////

	paramLookupExt    = "extension"
	paramLookupParams = "params"

	///////////////////////
	// method type param //
	///////////////////////

	category            = "CATEGORY"
	categoryChildren    = "CATEGORY_CHILDREN"
	categoryRelated     = "CATEGORY_RELATED"
	categorySeries      = "CATEGORY_SERIES"
	categoryTags        = "CATEGORY_TAGS"
	categoryRelatedTags = "CATEGORY_RELATED_TAGS"

	releases           = "RELEASES"
	releasesDates      = "RELEASES_DATES"
	release            = "RELEASE"
	releaseDates       = "RELEASE_DATES"
	releaseSeries      = "RELEASE_SERIES"
	releaseSources     = "RELEASE_SOURCES"
	releaseTags        = "RELEASE_TAGS"
	releaseRelatedTags = "RELEASE_RELATED_TAGS"
	releaseTables      = "RELEASE_TABLES"

	series                  = "SERIES"
	seriesCategories        = "SERIES_CATEGORIES"
	seriesObservations      = "SERIES_OBSERVATIONS"
	seriesRelease           = "SERIES_RELEASE"
	seriesSearch            = "SERIES_SEARCH"
	seriesSearchTags        = "SERIES_SEARCH_TAGS"
	seriesSearchRelatedTags = "SERIES_SEARCH_RELATED_TAGS"
	seriesTags              = "SERIES_TAGS"
	seriesUpdates           = "SERIES_UPDATES"
	seriesVintagedates      = "SERIES_VINTAGEDATES"

	sources        = "SOURCES"
	source         = "SOURCE"
	sourceReleases = "SOURCE_RELEASES"

	tags        = "TAGS"
	relatedTags = "RELATED_TAGS"
	tagsSeries  = "TAGS_SERIES"
)
