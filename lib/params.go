package lib

var paramsLookup = map[string]map[string]interface{}{
	// CATEGORY PARAMATERS

	"CATEGORY": {
		paramLookupExt:    "category",
		paramLookupParams: []string{"catetory_id"},
	},
	"CATEGORY_CHILDREN": {
		paramLookupExt:    "/category/children",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end"},
	},
	"CATEGORY_RELATED": {
		paramLookupExt:    "/category/related",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end"},
	},

	"CATEGORY_SERIES": {
		paramLookupExt:    "/category/series",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	"CATEGORY_TAGS": {
		paramLookupExt:    "/category/tags",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "tag_group_id", "search_text"},
	},
	"CATEGORY_RELATED_TAGS": {
		paramLookupExt:    "/category/tags",
		paramLookupParams: []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id", "search_text"},
	},

	// RELEASE PARAMATERS
	"RELEASES": {
		paramLookupExt:    "/releases",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	"RELEASES_DATES": {
		paramLookupExt:    "/releases/dates",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	"RELEASE": {
		paramLookupExt:    "/release",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end"},
	},
	"RELEASE_DATES": {
		paramLookupExt:    "/release/dates",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	"RELEASE_SERIES": {
		paramLookupExt:    "/release/series",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	"RELEASE_SOURCES": {
		paramLookupExt:    "/release/sources",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end"},
	},
	"RELEASE_TAGS": {
		paramLookupExt:    "/release/tags",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"RELEASE_RELATED_TAGS": {
		paramLookupExt:    "/release/related_tags",
		paramLookupParams: []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "exclude_tag_names", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"RELEASE_TABLES": {
		paramLookupExt:    "/release/tables",
		paramLookupParams: []string{"release_id", "element_id", "include_observation_values", "observation_date"},
	},

	// SERIES PARAMATERS
	"SERIES": {
		paramLookupExt:    "/series",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	"SERIES_CATEGORIES": {
		paramLookupExt:    "/series/categories",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	"SERIES_OBSERVATIONS": {
		paramLookupExt:    "/series/observations",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order", "observation_start", "observation_end", "units", "frequency", "aggregation_method", "output_type", "vintage_dates"},
	},
	"SERIES_RELEASE": {
		paramLookupExt:    "/series/release",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end"},
	},
	"SERIES_SEARCH": {
		paramLookupExt:    "/series/search",
		paramLookupParams: []string{"search_text", "search_type", "realtime_start", "realtime_end", "tag_names", "exclude_tag_names", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value"},
	},
	"SERIES_SEARCH_TAGS": {
		paramLookupExt:    "/series/search/tags",
		paramLookupParams: []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"SERIES_SEARCH_RELATED_TAGS": {
		paramLookupExt:    "/series/search/related_tags",
		paramLookupParams: []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"SERIES_TAGS": {
		paramLookupExt:    "/series/tags",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end", "order_by", "sort_order"},
	},
	"SERIES_UPDATES": {
		paramLookupExt:    "/series/updates",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "filter_value"},
	},
	"SERIES_VINTAGEDATES": {
		paramLookupExt:    "/series/vintagedates",
		paramLookupParams: []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order"},
	},

	// SOURCES PARAMATERS
	"SOURCES": {
		paramLookupExt:    "/sources",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	"SOURCE": {
		paramLookupExt:    "/source",
		paramLookupParams: []string{"source_id", "realtime_start", "realtime_end"},
	},
	"SOURCE_RELEASES": {
		paramLookupExt:    "/source/releases",
		paramLookupParams: []string{"source_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},

	// TAGS PARAMATERS
	"TAGS": {
		paramLookupExt:    "/tags",
		paramLookupParams: []string{"realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"RELATED_TAGS": {
		paramLookupExt:    "/related_tags",
		paramLookupParams: []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id"},
	},
	"TAGS_SERIES": {
		paramLookupExt:    "/tags/series",
		paramLookupParams: []string{"tag_names", "exclude_tag_names", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
}
