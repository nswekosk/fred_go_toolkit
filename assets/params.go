package lib

var ParamsLookup = map[string]map[string]interface{}{

	/////////////////////////
	// CATEGORY PARAMATERS //
	/////////////////////////

	CategoryParam: {
		ParamLookupExt:    "/category",
		ParamLookupParams: []string{"category_id"},
	},
	CategoryChildrenParam: {
		ParamLookupExt:    "/category/children",
		ParamLookupParams: []string{"category_id", "realtime_start", "realtime_end"},
	},
	CategoryRelatedParam: {
		ParamLookupExt:    "/category/related",
		ParamLookupParams: []string{"category_id", "realtime_start", "realtime_end"},
	},
	CategorySeriesParam: {
		ParamLookupExt:    "/category/series",
		ParamLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	CategoryTagsParam: {
		ParamLookupExt:    "/category/tags",
		ParamLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "tag_group_id", "search_text"},
	},
	CategoryRelatedTagsParam: {
		ParamLookupExt:    "/category/related_tags",
		ParamLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id", "search_text"},
	},

	////////////////////////
	// RELEASE PARAMATERS //
	////////////////////////

	ReleasesParam: {
		ParamLookupExt:    "/releases",
		ParamLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	ReleasesDatesParam: {
		ParamLookupExt:    "/releases/dates",
		ParamLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	ReleaseParam: {
		ParamLookupExt:    "/release",
		ParamLookupParams: []string{"release_id", "realtime_start", "realtime_end"},
	},
	ReleaseDatesParam: {
		ParamLookupExt:    "/release/dates",
		ParamLookupParams: []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	ReleaseSeriesParam: {
		ParamLookupExt:    "/release/series",
		ParamLookupParams: []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	ReleaseSourcesParam: {
		ParamLookupExt:    "/release/sources",
		ParamLookupParams: []string{"release_id", "realtime_start", "realtime_end"},
	},
	ReleaseTagsParam: {
		ParamLookupExt:    "/release/tags",
		ParamLookupParams: []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	ReleaseRelatedTagsParam: {
		ParamLookupExt:    "/release/related_tags",
		ParamLookupParams: []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "exclude_tag_names", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	ReleaseTablesParam: {
		ParamLookupExt:    "/release/tables",
		ParamLookupParams: []string{"release_id", "element_id", "include_observation_values", "observation_date"},
	},

	///////////////////////
	// SERIES PARAMATERS //
	///////////////////////

	SeriesParam: {
		ParamLookupExt:    "/series",
		ParamLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	SeriesCategoriesParam: {
		ParamLookupExt:    "/series/categories",
		ParamLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	SeriesObservationsParam: {
		ParamLookupExt:    "/series/observations",
		ParamLookupParams: []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order", "observation_start", "observation_end", "units", "frequency", "aggregation_method", "output_type", "vintage_dates"},
	},
	SeriesReleaseParam: {
		ParamLookupExt:    "/series/release",
		ParamLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	SeriesSearchParam: {
		ParamLookupExt:    "/series/search",
		ParamLookupParams: []string{"search_text", "search_type", "realtime_start", "realtime_end", "tag_names", "exclude_tag_names", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value"},
	},
	SeriesSearchTagsParam: {
		ParamLookupExt:    "/series/search/tags",
		ParamLookupParams: []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	SeriesSearchRelatedTagsParam: {
		ParamLookupExt:    "/series/search/related_tags",
		ParamLookupParams: []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	SeriesTagsParam: {
		ParamLookupExt:    "/series/tags",
		ParamLookupParams: []string{"series_id", "realtime_start", "realtime_end", "order_by", "sort_order"},
	},
	SeriesUpdatesParam: {
		ParamLookupExt:    "/series/updates",
		ParamLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "filter_value"},
	},
	SeriesVintagedatesParam: {
		ParamLookupExt:    "/series/vintagedates",
		ParamLookupParams: []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order"},
	},

	////////////////////////
	// SOURCES PARAMATERS //
	////////////////////////

	SourcesParam: {
		ParamLookupExt:    "/sources",
		ParamLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	SourceParam: {
		ParamLookupExt:    "/source",
		ParamLookupParams: []string{"source_id", "realtime_start", "realtime_end"},
	},
	SourceReleasesParam: {
		ParamLookupExt:    "/source/releases",
		ParamLookupParams: []string{"source_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},

	/////////////////////
	// TAGS PARAMATERS //
	/////////////////////

	TagsParam: {
		ParamLookupExt:    "/tags",
		ParamLookupParams: []string{"realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	RelatedTagsParam: {
		ParamLookupExt:    "/related_tags",
		ParamLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id"},
	},
	TagsSeriesParam: {
		ParamLookupExt:    "/tags/series",
		ParamLookupParams: []string{"tag_names", "exclude_tag_names", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
}
