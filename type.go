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
	ItemInfo SkinInfo `json:"iteminfo"`
}

type SkinInfo struct {
	Float    float64  `json:"floatvalue"`
	Stickers Stickers `json:"stickers"`
}

type Stickers []struct {
	Name string `json:"name"`
}

func (stickers Stickers) Format() string {
	formatedStickers := ""
	for index, sticker := range stickers {
		if index != 0 {
			formatedStickers += ", "
		}
		formatedStickers += sticker.Name
	}
	return formatedStickers
}
