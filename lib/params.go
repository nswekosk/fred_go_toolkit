package lib

var paramsLookup = map[string]map[string]interface{}{

	/////////////////////////
	// CATEGORY PARAMATERS //
	/////////////////////////

	categoryParam: {
		paramLookupExt:    "/category",
		paramLookupParams: []string{"category_id"},
	},
	categoryChildrenParam: {
		paramLookupExt:    "/category/children",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end"},
	},
	categoryRelatedParam: {
		paramLookupExt:    "/category/related",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end"},
	},
	categorySeriesParam: {
		paramLookupExt:    "/category/series",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	categoryTagsParam: {
		paramLookupExt:    "/category/tags",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "tag_group_id", "search_text"},
	},
	categoryRelatedTagsParam: {
		paramLookupExt:    "/category/related_tags",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id", "search_text"},
	},

	////////////////////////
	// RELEASE PARAMATERS //
	////////////////////////

	releasesParam: {
		paramLookupExt:    "/releases",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	releasesDatesParam: {
		paramLookupExt:    "/releases/dates",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	releaseParam: {
		paramLookupExt:    "/release",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end"},
	},
	releaseDatesParam: {
		paramLookupExt:    "/release/dates",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	releaseSeriesParam: {
		paramLookupExt:    "/release/series",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	releaseSourcesParam: {
		paramLookupExt:    "/release/sources",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end"},
	},
	releaseTagsParam: {
		paramLookupExt:    "/release/tags",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	releaseRelatedTagsParam: {
		paramLookupExt:    "/release/related_tags",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "exclude_tag_names", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	releaseTablesParam: {
		paramLookupExt:    "/release/tables",
		paramLookupParams: []string{"release_id", "element_id", "include_observation_values", "observation_date"},
	},

	///////////////////////
	// SERIES PARAMATERS //
	///////////////////////

	seriesParam: {
		paramLookupExt:    "/series",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	seriesCategoriesParam: {
		paramLookupExt:    "/series/categories",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	seriesObservationsParam: {
		paramLookupExt:    "/series/observations",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order", "observation_start", "observation_end", "units", "frequency", "aggregation_method", "output_type", "vintage_dates"},
	},
	seriesReleaseParam: {
		paramLookupExt:    "/series/release",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	seriesSearchParam: {
		paramLookupExt:    "/series/search",
		paramLookupParams: []string{"search_text", "search_type", "realtime_start", "realtime_end", "tag_names", "exclude_tag_names", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value"},
	},
	seriesSearchTagsParam: {
		paramLookupExt:    "/series/search/tags",
		paramLookupParams: []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	seriesSearchRelatedTagsParam: {
		paramLookupExt:    "/series/search/related_tags",
		paramLookupParams: []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	seriesTagsParam: {
		paramLookupExt:    "/series/tags",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end", "order_by", "sort_order"},
	},
	seriesUpdatesParam: {
		paramLookupExt:    "/series/updates",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "filter_value"},
	},
	seriesVintagedatesParam: {
		paramLookupExt:    "/series/vintagedates",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order"},
	},

	////////////////////////
	// SOURCES PARAMATERS //
	////////////////////////

	sourcesParam: {
		paramLookupExt:    "/sources",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	sourceParam: {
		paramLookupExt:    "/source",
		paramLookupParams: []string{"source_id", "realtime_start", "realtime_end"},
	},
	sourceReleasesParam: {
		paramLookupExt:    "/source/releases",
		paramLookupParams: []string{"source_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},

	/////////////////////
	// TAGS PARAMATERS //
	/////////////////////

	tagsParam: {
		paramLookupExt:    "/tags",
		paramLookupParams: []string{"realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	relatedTagsParam: {
		paramLookupExt:    "/related_tags",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id"},
	},
	tagsSeriesParam: {
		paramLookupExt:    "/tags/series",
		paramLookupParams: []string{"tag_names", "exclude_tag_names", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
}
