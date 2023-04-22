package hamrobazar

const (
	CategoryHouse = "5c236040-eab5-4c16-8954-41a0f2f780ce"
)

type SortBy int

const (
	SortLatest SortBy = iota
)

type Config struct {
	Label    string `yaml:"label"`
	Disabled bool   `yaml:"disabled"`
	Filter   Filter `yaml:"filter"`
}

type Filter struct {
	PageNumber   int          `yaml:"page_number"`
	PageSize     int          `yaml:"page_size"`
	Latitude     float64      `yaml:"latitude"`
	Longitude    float64      `yaml:"longitude"`
	SearchParams SearchParams `yaml:"search_params"`
	FilterParams FilterParams `yaml:"filter_params"`
	SortParam    SortBy       `yaml:"sort_param"`
	DeviceId     string       `yaml:"device_id"`
	DeviceSource string       `yaml:"device_source"`
}

type SearchParams struct {
	Value    string `yaml:"searchValue"`
	SearchBy string `yaml:"search_by"`
}

type FilterParams struct {
	Condition         int    `yaml:"condition"`
	PriceFrom         int    `yaml:"price_from"`
	PriceTo           int    `yaml:"price_to"`
	IsPriceNegotiable *bool  `yaml:"is_price_negotiable"`
	Category          string `yaml:"category"`
	CategoryIds       string `yaml:"category_ids"`
	BrandIds          string `yaml:"brand_ids"`
	Brand             string `yaml:"brand"`
}

type SearchResult struct {
	Data []Item
}

type Item struct {
	ID          string
	Name        string
	Price       float64
	ImageURL    string
	CreatorInfo Creator
	Description string
	Location    Location
}

type Creator struct {
	CreatedByName  string
	CreatorAdCount int
}

type Location struct {
	LocationDescription string
}
