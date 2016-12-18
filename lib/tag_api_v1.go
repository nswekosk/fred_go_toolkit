package lib

type Tags struct {
	Start     string   `realtime_start`
	End       string   `realtime_end`
	OrderBy   string   `order_by`
	SortOrder string   `sort_order`
	Count     int      `count`
	Offset    int      `offset`
	Limit     int      `limit`
	Tags      []Tag    `tag`
	Sources   []Source `sources`
}

type Tag struct {
	Name        string `name`
	GroupId     string `group_id`
	Notes       string `notes`
	Created     string `created`
	Popularity  int    `popularity`
	SeriesCount int    `series_count`
}
