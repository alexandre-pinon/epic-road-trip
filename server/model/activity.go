package model

type Activity struct {
	HTMLAttributions []interface{}    `json:"html_attributions"`
	Results          []ActivityResult `json:"results"`
	Status           string           `json:"status"`
}

type ActivityResult struct {
	BusinessStatus      string           `json:"business_status"`
	Geometry            GeometryActivity `json:"geometry"`
	Icon                string           `json:"icon"`
	IconBackgroundColor string           `json:"icon_background_color"`
	IconMaskBaseURI     string           `json:"icon_mask_base_uri"`
	Name                string           `json:"name"`
	OpeningHours        OpeningHours     `json:"opening_hours,omitempty"`
	Photos              []Photos         `json:"photos"`
	PlaceID             string           `json:"place_id"`
	PlusCode            PlusCode         `json:"plus_code"`
	PriceLevel          int              `json:"price_level,omitempty"`
	Rating              float64          `json:"rating"`
	Reference           string           `json:"reference"`
	Scope               string           `json:"scope"`
	Types               []string         `json:"types"`
	UserRatingsTotal    int              `json:"user_ratings_total"`
	Vicinity            string           `json:"vicinity"`
}

type GeometryActivity struct {
	Location Location `json:"location"`
	Viewport Bounds   `json:"viewport"`
}

type OpeningHours struct {
	OpenNow bool `json:"open_now"`
}

type Photos struct {
	Height           int      `json:"height"`
	HTMLAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int      `json:"width"`
}

type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}
