package lib

var paramsLookup = map[string][]string{
	// CATEGORY PARAMATERS
	"CATEGORY":              []string{"catetory_id"},
	"CATEGORY_CHILDREN":     []string{"category_id", "realtime_start", "realtime_end"},
	"CATEGORY_RELATED":      []string{"category_id", "realtime_start", "realtime_end"},
	"CATEGORY_SERIES":       []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	"CATEGORY_TAGS":         []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "tag_group_id", "search_text"},
	"CATEGORY_RELATED_TAGS": []string{"category_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id", "search_text"},

	// RELEASE PARAMATERS
	"RELEASES":             []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	"RELEASES_DATES":       []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	"RELEASE":              []string{"release_id", "realtime_start", "realtime_end"},
	"RELEASE_DATES":        []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "include_release_dates_with_no_data"},
	"RELEASE_SERIES":       []string{"release_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value", "tag_names", "exclude_tag_names"},
	"RELEASE_SOURCES":      []string{"release_id", "realtime_start", "realtime_end"},
	"RELEASE_TAGS":         []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	"RELEASE_RELATED_TAGS": []string{"release_id", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "exclude_tag_names", "search_text", "limit", "offset", "order_by", "sort_order"},
	"RELEASE_TABLES":       []string{"release_id", "element_id", "include_observation_values", "observation_date"},

	// SERIES PARAMATERS
	"SERIES":                     []string{"series_id", "realtime_start", "realtime_end"},
	"SERIES_CATEGORIES":          []string{"series_id", "realtime_start", "realtime_end"},
	"SERIES_OBSERVATIONS":        []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order", "observation_start", "observation_end", "units", "frequency", "aggregation_method", "output_type", "vintage_dates"},
	"SERIES_RELEASE":             []string{"series_id", "realtime_start", "realtime_end"},
	"SERIES_SEARCH":              []string{"search_text", "search_type", "realtime_start", "realtime_end", "tag_names", "exclude_tag_names", "limit", "offset", "order_by", "sort_order", "filter_variable", "filter_value"},
	"SERIES_SEARCH_TAGS":         []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	"SERIES_SEARCH_RELATED_TAGS": []string{"series_search_text", "realtime_start", "realtime_end", "tag_names", "tag_group_id", "tag_search_text", "limit", "offset", "order_by", "sort_order"},
	"SERIES_TAGS":                []string{"series_id", "realtime_start", "realtime_end", "order_by", "sort_order"},
	"SERIES_UPDATES":             []string{"realtime_start", "realtime_end", "limit", "offset", "filter_value"},
	"SERIES_VINTAGEDATES":        []string{"series_id", "realtime_start", "realtime_end", "limit", "offset", "sort_order"},

	// SOURCES PARAMATERS
	"SOURCES":         []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
	"SOURCE":          []string{"source_id", "realtime_start", "realtime_end"},
	"SOURCE_RELEASES": []string{"source_id", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},

	// TAGS PARAMATERS
	"TAGS":         []string{"realtime_start", "realtime_end", "tag_names", "tag_group_id", "search_text", "limit", "offset", "order_by", "sort_order"},
	"RELATED_TAGS": []string{"realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order", "tag_names", "exclude_tag_names", "tag_group_id"},
	"TAGS_SERIES":  []string{"tag_names", "exclude_tag_names", "realtime_start", "realtime_end", "limit", "offset", "order_by", "sort_order"},
}
