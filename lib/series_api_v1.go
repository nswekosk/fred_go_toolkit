package lib

type Seriess struct {
	Start     string   `realtime_start`
	End       string   `realtime_end`
	SeriesCol []Series `seriess`
}

type Series struct {
	ID                     string `id`
	Start                  string `realtime_start`
	End                    string `realtime_end`
	Title                  string `title`
	ObsStart               string `observation_start`
	ObsEnd                 string `observation_end`
	Frequency              string `annual`
	FrequencyShort         string `frequency_short`
	Units                  string `units`
	UnitsShort             string `units_short`
	SeasonalAdjustment     string `seasonal_adjustment`
	SeasonalAdustmentShort string `seasonal_adjustment_short`
	LastUpdated            string `last_updated`
	Popularity             int    `popularity`
	Notes                  string `notes`
}
