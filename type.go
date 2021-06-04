package main

type MarketListingItem struct {
	ListingID      string `json:"listingid"`
	ConvertedPrice int    `json:"converted_price"`
	ConvertedFee   int    `json:"converted_fee"`
	Asset          struct {
		ID       string `json:"id"`
		Acctions []struct {
			InspectLink string `json:"link"`
		} `json:"market_actions"`
	} `json:"asset"`
}

func (m MarketListingItem) Price() float64 {
	return float64(m.ConvertedPrice+m.ConvertedFee) * 0.01
}

type MarketListingResponse struct {
	Items map[string]MarketListingItem `json:"listinginfo"`
}

type GetFloatResponse struct {
	Info struct {
		Float float64 `json:"floatvalue"`
	} `json:"iteminfo"`
}
