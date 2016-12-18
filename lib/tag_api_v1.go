package lib

type Tags struct {
	Start     string   `realtime_start`
	End       string   `realtime_end`
	OrderBy   string   `order_by`
	SortOrder string   `sort_order`
	Count     int      `count`
	Offset    int      `offset`
	Limit     int      `limit`
	Sources   []Source `sources`
}
