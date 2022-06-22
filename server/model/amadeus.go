package model

import "time"

type AccessTokenResponse struct {
	Type            string `json:"type"`
	Username        string `json:"username"`
	ApplicationName string `json:"application_name"`
	ClientID        string `json:"client_id"`
	TokenType       string `json:"token_type"`
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	State           string `json:"state"`
	Scope           string `json:"scope"`
}

type AccessTokenError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Code             int    `json:"code"`
	Title            string `json:"title"`
}

type FlightFormData struct {
	OriginLocationCode      string    `json:"origin_location_code" binding:"required"`
	DestinationLocationCode string    `json:"destination_location_code" binding:"required"`
	DepartureDate           time.Time `json:"departure_date" binding:"required"`
	Adults                  int       `json:"adults" binding:"required"`
}

type FlighOffersResponse struct {
	Meta         FlightOfferMeta       `json:"meta"`
	Data         []FlightOffer         `json:"data"`
	Dictionaries FlightOfferDictionary `json:"dictionaries,omitempty"`
}

type FlightOfferMeta struct {
	Count int                  `json:"count"`
	Links FlightOfferMetaLinks `json:"links"`
}

type FlightOfferMetaLinks struct {
	Self string `json:"self"`
}

type FlightOffer struct {
	Type                     string            `json:"type"`
	ID                       string            `json:"id"`
	Source                   string            `json:"source"`
	InstantTicketingRequired bool              `json:"instant_ticketing_required"`
	NonHomogeneous           bool              `json:"non_homogeneous"`
	OneWay                   bool              `json:"one_way"`
	LastTicketingDate        string            `json:"last_ticketing_date"`
	NumberOfBookableSeats    int               `json:"number_of_bookable_seats"`
	Itineraries              []Itinary         `json:"itinaries"`
	Price                    Price             `json:"price"`
	PricingOptions           PricingOptions    `json:"pricing_options"`
	ValidatingAirlineCodes   []string          `json:"validating_airline_codes"`
	TravelerPricings         []TravelerPricing `json:"travel_pricings"`
}

type Itinary struct {
	Duration string    `json:"duration"`
	Segments []Segment `json:"segments"`
}

type Segment struct {
	Departure       Departure `json:"departure"`
	Arrival         Arrival   `json:"arrival"`
	CarrierCode     string    `json:"carrier_code"`
	Number          string    `json:"number"`
	Aircraft        Aircraft  `json:"aircraft"`
	Operating       Operating `json:"operating"`
	Duration        string    `json:"duration"`
	ID              string    `json:"id"`
	NumberOfStops   int       `json:"number_of_stops"`
	BlacklistedInEU bool      `json:"blacklisted_in_EU"`
}

type Departure struct {
	IataCode string `json:"iata_code"`
	Terminal string `json:"terminal"`
	At       string `json:"at"`
}

type Arrival struct {
	IataCode string `json:"iata_code"`
	At       string `json:"at"`
}

type Aircraft struct {
	Code string `json:"code"`
}

type Operating struct {
	CarrierCode string `json:"carrier_code"`
}

type BasePrice struct {
	Currency string `json:"currency"`
	Total    string `json:"total"`
	Base     string `json:"base"`
}

type Price struct {
	BasePrice
	Fees       []Fee  `json:"fees"`
	GrandTotal string `json:"grand_total"`
}

type Fee struct {
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type PricingOptions struct {
	FareType                []string `json:"fare_type"`
	IncludedCheckedBagsOnly bool     `json:"included_checked_bags_only"`
}

type TravelerPricing struct {
	TravelerID           string        `json:"traveler_id"`
	FareOption           string        `json:"fare_option"`
	TravelerType         string        `json:"traveler_type"`
	Price                BasePrice     `json:"price"`
	FareDetailsBySegment []FareDetails `json:"fare_details_by_segment"`
}

type FareDetails struct {
	SegmentID           string              `json:"segmentId"`
	Cabin               string              `json:"cabin"`
	FareBasis           string              `json:"fareBasis"`
	Class               string              `json:"class"`
	IncludedCheckedBags IncludedCheckedBags `json:"included_checked_bags"`
}

type IncludedCheckedBags struct {
	Quantity int `json:"quantity"`
}

type FlightOfferDictionary struct {
	Locations  map[string]FlightOfferLocations `json:"locations"`
	Aircraft   map[string]string               `json:"aircraft"`
	Currencies map[string]string               `json:"currencies"`
	Carriers   map[string]string               `json:"carriers"`
}

type FlightOfferLocations struct {
	CityCode    string `json:"cityCode"`
	CountryCode string `json:"countryCode"`
}
