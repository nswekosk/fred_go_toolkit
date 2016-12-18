package lib

type Sources struct {
	Start      string   `realtime_start`
	End        string   `realtime_end`
	OrderBy    string   `order_by`
	SortOrder  string   `sort_order`
	Count      int      `count`
	Offset     int      `offset`
	Limit      int      `limit`
	SourcesCol []Source `sources`
}

type Source struct {
	ID    int    `id`
	Start string `realtime_start`
	End   string `realtime_end`
	Name  string `name`
	Link  string `link`
}
