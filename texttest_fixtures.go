package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/myugen/go-gildedrose-solution/catalogue"
	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/myugen/go-gildedrose-solution/item"
)

func main() {
	fmt.Println("OMGHAI!")

	var valuableItems = []catalogue.ValuableItem{
		catalogue.NewRegularValuableItem(&item.Item{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20}).ValuableItem,
		catalogue.NewAgedValuableItem(&item.Item{Name: "Aged Brie", SellIn: 2, Quality: 0}).ValuableItem,
		catalogue.NewRegularValuableItem(&item.Item{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7}).ValuableItem,
		catalogue.NewLegendaryValuableItem(&item.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80}).ValuableItem,
		catalogue.NewLegendaryValuableItem(&item.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80}).ValuableItem,
		catalogue.NewEventualValuableItem(&item.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20}).ValuableItem,
		catalogue.NewEventualValuableItem(&item.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 49}).ValuableItem,
		catalogue.NewEventualValuableItem(&item.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 49}).ValuableItem,
		catalogue.NewRegularValuableItem(&item.Item{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6}).ValuableItem,
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("name, sellIn, quality")
		for i := 0; i < len(valuableItems); i++ {
			fmt.Println(valuableItems[i])
		}
		fmt.Println("")
		gildedrose.UpdateQuality(valuableItems...)
	}
}
