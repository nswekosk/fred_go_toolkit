package lib

var paramsLookup = map[string]map[string]interface{}{
	// CATEGORY PARAMATERS

	"CATEGORY": {
		"extension": "category",
		"params":    []string{"catetory_id"},
	},
	"CATEGORY_CHILDREN": {
		"extension": "/category/children",
		"params":    []string{"category_id", "realtime_start", "realtime_end"},
	},
	"CATEGORY_RELATED": {
		"extension": "/category/related",
		"params":    []string{"category_id", "realtime_start", "realtime_end"},
	},

	"CATEGORY_SERIES": {
		"extension": "/category/series",
		"params":    []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	"CATEGORY_TAGS": {
		"extension": "/category/tags",
		"params":    []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "tag_group_id", "search_text"},
	},
	"CATEGORY_RELATED_TAGS": {
		"extension": "/category/tags",
		"params":    []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id", "search_text"},
	},

	// RELEASE PARAMATERS
	"RELEASES": {
		"extension": "/releases",
		"params":    []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	"RELEASES_DATES": {
		"extension": "/releases/dates",
		"params":    []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	"RELEASE": {
		"extension": "/release",
		"params":    []string{"release_id", "realtime_start", "realtime_end"},
	},
	"RELEASE_DATES": {
		"extension": "/release/dates",
		"params":    []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	},
	"RELEASE_SERIES": {
		"extension": "/release/series",
		"params":    []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	},
	"RELEASE_SOURCES": {
		"extension": "/release/sources",
		"params":    []string{"release_id", "realtime_start", "realtime_end"},
	},
	"RELEASE_TAGS": {
		"extension": "/release/tags",
		"params":    []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"RELEASE_RELATED_TAGS": {
		"extension": "/release/related_tags",
		"params":    []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "exclude_tag_names", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"RELEASE_TABLES": {
		"extension": "/release/tables",
		"params":    []string{"release_id", "element_id", "include_observation_values", "observation_date"},
	},

	// SERIES PARAMATERS
	"SERIES": {
		"extension": "/series",
		"params":    []string{"series_id", "realtime_start", "realtime_end"},
	},
	"SERIES_CATEGORIES": {
		"extension": "/series/categories",
		"params":    []string{"series_id", "realtime_start", "realtime_end"},
	},
	"SERIES_OBSERVATIONS": {
		"extension": "/series/observations",
		"params":    []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order", "observation_start", "observation_end", "units", "frequency", "aggregation_method", "output_type", "vintage_dates"},
	},
	"SERIES_RELEASE": {
		"extension": "/series/release",
		"params":    []string{"series_id", "realtime_start", "realtime_end"},
	},
	"SERIES_SEARCH": {
		"extension": "/series/search",
		"params":    []string{"search_text", "search_type", "realtime_start", "realtime_end", "tag_names", "exclude_tag_names", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value"},
	},
	"SERIES_SEARCH_TAGS": {
		"extension": "/series/search/tags",
		"params":    []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"SERIES_SEARCH_RELATED_TAGS": {
		"extension": "/series/search/related_tags",
		"params":    []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"SERIES_TAGS": {
		"extension": "/series/tags",
		"params":    []string{"series_id", "realtime_start", "realtime_end", "order_by", "sort_order"},
	},
	"SERIES_UPDATES": {
		"extension": "/series/updates",
		"params":    []string{"realtime_start", "realtime_end", "limit", "offset", "filter_value"},
	},
	"SERIES_VINTAGEDATES": {
		"extension": "/series/vintagedates",
		"params":    []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order"},
	},

	// SOURCES PARAMATERS
	"SOURCES": {
		"extension": "/sources",
		"params":    []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
	"SOURCE": {
		"extension": "/source",
		"params":    []string{"source_id", "realtime_start", "realtime_end"},
	},
	"SOURCE_RELEASES": {
		"extension": "/source/releases",
		"params":    []string{"source_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},

	// TAGS PARAMATERS
	"TAGS": {
		"extension": "/tags",
		"params":    []string{"realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	},
	"RELATED_TAGS": {
		"extension": "/related_tags",
		"params":    []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id"},
	},
	"TAGS_SERIES": {
		"extension": "/tags/series",
		"params":    []string{"tag_names", "exclude_tag_names", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	},
}
