package lib

var paramsLookup = map[string]map[string]interface{}{

	/////////////////////////
	// CATEGORY PARAMATERS //
	/////////////////////////

	category: {
		paramLookupExt:    "category",
		paramLookupParams: []string{"catetory_id"},
	},
	categoryChildren: {
		paramLookupExt:    "/category/children",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end"},
	},
	categoryRelated: {
		paramLookupExt:    "/category/related",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end"},
	},
	categorySeries: {
		paramLookupExt:    "/category/series",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	categoryTags: {
		paramLookupExt:    "/category/tags",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "tag_group_id", "search_text"},
	},
	categoryRelatedTags: {
		paramLookupExt:    "/category/tags",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id", "search_text"},
	},

	////////////////////////
	// RELEASE PARAMATERS //
	////////////////////////

	releases: {
		paramLookupExt:    "/releases",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	releasesDates: {
		paramLookupExt:    "/releases/dates",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	release: {
		paramLookupExt:    "/release",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end"},
	},
	releaseDates: {
		paramLookupExt:    "/release/dates",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	releaseSeries: {
		paramLookupExt:    "/release/series",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	releaseSources: {
		paramLookupExt:    "/release/sources",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end"},
	},
	releaseTags: {
		paramLookupExt:    "/release/tags",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	releaseRelatedTags: {
		paramLookupExt:    "/release/related_tags",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "exclude_tag_names", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	releaseTables: {
		paramLookupExt:    "/release/tables",
		paramLookupParams: []string{"release_id", "element_id", "include_observation_values", "observation_date"},
	},

	///////////////////////
	// SERIES PARAMATERS //
	///////////////////////

	series: {
		paramLookupExt:    "/series",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	seriesCategories: {
		paramLookupExt:    "/series/categories",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	seriesObservations: {
		paramLookupExt:    "/series/observations",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order", "observation_start", "observation_end", "units", "frequency", "aggregation_method", "output_type", "vintage_dates"},
	},
	seriesRelease: {
		paramLookupExt:    "/series/release",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	seriesSearch: {
		paramLookupExt:    "/series/search",
		paramLookupParams: []string{"search_text", "search_type", "realtime_start", "realtime_end", "tag_names", "exclude_tag_names", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value"},
	},
	seriesSearchTags: {
		paramLookupExt:    "/series/search/tags",
		paramLookupParams: []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	seriesSearchRelatedTags: {
		paramLookupExt:    "/series/search/related_tags",
		paramLookupParams: []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	seriesTags: {
		paramLookupExt:    "/series/tags",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end", "order_by", "sort_order"},
	},
	seriesUpdates: {
		paramLookupExt:    "/series/updates",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "filter_value"},
	},
	seriesVintagedates: {
		paramLookupExt:    "/series/vintagedates",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order"},
	},

	////////////////////////
	// SOURCES PARAMATERS //
	////////////////////////

	sources: {
		paramLookupExt:    "/sources",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	source: {
		paramLookupExt:    "/source",
		paramLookupParams: []string{"source_id", "realtime_start", "realtime_end"},
	},
	sourceReleases: {
		paramLookupExt:    "/source/releases",
		paramLookupParams: []string{"source_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},

	/////////////////////
	// TAGS PARAMATERS //
	/////////////////////

	tags: {
		paramLookupExt:    "/tags",
		paramLookupParams: []string{"realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	relatedTags: {
		paramLookupExt:    "/related_tags",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id"},
	},
	tagsSeries: {
		paramLookupExt:    "/tags/series",
		paramLookupParams: []string{"tag_names", "exclude_tag_names", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
}
