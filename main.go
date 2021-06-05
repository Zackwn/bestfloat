package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
)

func GetItemBuyCommand(item *MarketListingItem) string {
	return fmt.Sprintf("BuyMarketListing('listing', '%v', 730, '2', '%v')", item.ListingID, item.Asset.ID)
}

func DisplayItem(item *MarketListingItem, info *SkinInfo, tab int) {
	fmt.Println("Tab:", tab)
	fmt.Println("Float: ", info.Float)
	fmt.Println(len(info.Stickers), "Stickers:", info.Stickers.Format())
	fmt.Printf("Price: %v$\n", item.Price())
	fmt.Println("Buying Command:", GetItemBuyCommand(item))
	fmt.Println("-------------------------------------------------------------------------------------------")
}

func main() {
	fmt.Print("Skin hash name: ")

	skinName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	skinName = skinName[:len(skinName)-2] // remove "\n"
	skinName = strings.ReplaceAll(url.QueryEscape(skinName), "+", "%20")

	fmt.Println("-------------------------------------------------------------------------------------------")

	var wg sync.WaitGroup
	const maxRequests = 20
	end := make(chan bool)
	i, start, count, tab := 0, 0, 10, 1
	for i < maxRequests {
		wg.Add(1)
		go func(start, count, tab int) {
			defer wg.Done()
			listings, err := GetSteamListings(start, count, skinName)
			if err != nil {
				end <- true
				return
			}
			end <- false
			bestFloatInfo := new(SkinInfo)
			bestFloatInfo.Float = 1
			var bestFloatItem MarketListingItem
			for _, item := range listings.Items {
				skinInfo := GetSkinInfo(&item)
				if skinInfo.Float < bestFloatInfo.Float {
					bestFloatItem = item
					bestFloatInfo = skinInfo
				}
			}
			DisplayItem(&bestFloatItem, bestFloatInfo, tab)
		}(start, count, tab)

		if <-end {
			break
		}
		start += count
		i++
		tab++
	}

	wg.Wait()
}
