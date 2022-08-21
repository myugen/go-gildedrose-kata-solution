package gildedrose

import "github.com/myugen/go-gildedrose-kata-solution/catalogue"

func UpdateQuality(items ...catalogue.ValuableItem) {
	for _, item := range items {
		item.Update()
	}
}
