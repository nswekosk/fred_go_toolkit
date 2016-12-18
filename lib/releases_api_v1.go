package lib

type Releases struct {
	Start      string    `realtime_start`
	End        string    `realtime_end`
	OrderBy    string    `order_by`
	SortOrder  string    `sort_order`
	Count      int       `count`
	Offset     int       `offset`
	Limit      int       `limit`
	ReleaseCol []Release `releases`
}

type Release struct {
	ID           int    `id`
	Start        string `realtime_start`
	End          string `realtime_end`
	Name         string `name`
	PressRelease bool   `press_release`
	Link         string `link`
}
