package model

import "time"

type AccessTokenResponse struct {
	Type            string `json:"type"`
	Username        string `json:"username"`
	ApplicationName string `json:"application_name"`
	ClientID        string `json:"clien_id"`
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

type AccessToken struct {
	Value string `json:"value"`
	Iat   int64  `json:"iat"`
	Exp   int    `json:"exp"`
}

type FlightFormData struct {
	OriginLocation          string    `json:"originLocation" binding:"required"`
	DestinationLocation     string    `json:"destinationLocation" binding:"required"`
	OriginLocationCode      string    `json:"originLocationCode"`
	DestinationLocationCode string    `json:"destinationLocationCode"`
	DepartureDate           time.Time `json:"departureDate" binding:"required"`
	Adults                  int       `json:"adults" binding:"required"`
}

type FlightOffersResponse struct {
	Meta         FlightOfferMeta       `json:"meta"`
	Data         []FlightOffer         `json:"data"`
	Dictionaries FlightOfferDictionary `json:"dictionaries,omitempty"`
}

type FlightOffersResponseError struct {
	Errors FlightOffersErrors `json:"errors"`
}

type FlightOffersErrors struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

type FlightOfferMeta struct {
	Count int                  `json:"count"`
	Links FlightOfferMetaLinks `json:"links"`
}

type FlightOfferMetaLinks struct {
	Self string `json:"self"`
}

type FlightOffer struct {
	Type                     string               `json:"type"`
	ID                       string               `json:"id"`
	Source                   string               `json:"source"`
	InstantTicketingRequired bool                 `json:"instantTicketingRequired"`
	NonHomogeneous           bool                 `json:"nonHomogeneous"`
	OneWay                   bool                 `json:"oneWay"`
	LastTicketingDate        string               `json:"lastTicketingDate"`
	NumberOfBookableSeats    int                  `json:"numberOfBookableSeats"`
	Itineraries              []FlightOfferItinary `json:"itineraries"`
	Price                    Price                `json:"price"`
	PricingOptions           PricingOptions       `json:"pricingOptions"`
	ValidatingAirlineCodes   []string             `json:"validatingAirlineCodes"`
	TravelerPricings         []TravelerPricing    `json:"travelPricings"`
}

type FlightOfferItinary struct {
	Duration string    `json:"duration"`
	Segments []Segment `json:"segments"`
}

type Segment struct {
	Departure       Departure `json:"departure"`
	Arrival         Arrival   `json:"arrival"`
	CarrierCode     string    `json:"carrierCode"`
	Number          string    `json:"number"`
	Aircraft        Aircraft  `json:"aircraft"`
	Operating       Operating `json:"operating"`
	Duration        string    `json:"duration"`
	ID              string    `json:"id"`
	NumberOfStops   int       `json:"numberOfStops"`
	BlacklistedInEU bool      `json:"blacklistedInEU"`
}

type Departure struct {
	IataCode string `json:"iataCode"`
	Terminal string `json:"terminal"`
	At       string `json:"at"`
}

type Arrival struct {
	IataCode string `json:"iataCode"`
	At       string `json:"at"`
}

type Aircraft struct {
	Code string `json:"code"`
}

type Operating struct {
	CarrierCode string `json:"carrierCode"`
}

type BasePrice struct {
	Currency string `json:"currency"`
	Total    string `json:"total"`
	Base     string `json:"base"`
}

type Price struct {
	BasePrice
	Fees       []Fee  `json:"fees"`
	GrandTotal string `json:"grandTotal"`
}

type Fee struct {
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type PricingOptions struct {
	FareType                []string `json:"fareType"`
	IncludedCheckedBagsOnly bool     `json:"includedCheckedBagsOnly"`
}

type TravelerPricing struct {
	TravelerID           string        `json:"travelerId"`
	FareOption           string        `json:"fareOption"`
	TravelerType         string        `json:"travelerType"`
	Price                BasePrice     `json:"price"`
	FareDetailsBySegment []FareDetails `json:"fareDetailsBySegment"`
}

type FareDetails struct {
	SegmentID           string              `json:"segmentId"`
	Cabin               string              `json:"cabin"`
	FareBasis           string              `json:"fareBasis"`
	Class               string              `json:"class"`
	IncludedCheckedBags IncludedCheckedBags `json:"includedCheckedBags"`
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
