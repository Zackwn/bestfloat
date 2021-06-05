package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetSteamListings(start, count int, skinName string) (*MarketListingResponse, error) {
	url := fmt.Sprintf("https://steamcommunity.com/market/listings/730/%v/render?start=%v&count=%v&currency=1&language=english&norender=1&format=json", skinName, start, count)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	} else if response.StatusCode != 200 {
		log.Fatalln(errors.New(response.Status))
	}

	defer response.Body.Close()
	listings := new(MarketListingResponse)
	json.NewDecoder(response.Body).Decode(listings)
	if len(listings.Items) == 0 {
		return nil, errors.New("GetSteamListings: no more items")
	}
	return listings, nil
}

func GetSkinInfo(item *MarketListingItem) *SkinInfo {
	inspectLink := strings.Replace(item.Asset.Acctions[0].InspectLink, "%listingid%", item.ListingID, 1)
	inspectLink = strings.Replace(inspectLink, "%assetid%", item.Asset.ID, 1)

	response, err := http.Get("https://floats.gainskins.com/?url=" + inspectLink)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	res := new(GetFloatResponse)
	json.NewDecoder(response.Body).Decode(res)
	return &res.ItemInfo
}
