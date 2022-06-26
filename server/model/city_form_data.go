package model

type CityFormData struct {
	City string `json:"city"`
	Constraints Constraints `json:"constraints"`
}

type Constraints struct {
	Radius		int 	`json:"radius"`
	MaxPrice 	int 	`json:"maxprice"`
	MinPrice	int 	`json:"minprice"`
	OpenNow		bool 	`json:"opennow"`
}