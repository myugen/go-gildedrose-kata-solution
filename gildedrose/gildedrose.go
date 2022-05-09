package gildedrose

import "github.com/myugen/go-gildedrose-solution/catalogue"

func UpdateQuality(items ...catalogue.ValuableItem) {
	for _, item := range items {
		item.Update()
	}
}
