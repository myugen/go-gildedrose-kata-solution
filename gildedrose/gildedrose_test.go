package gildedrose_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/stretchr/testify/assert"
)

func TestUpdateQuality(t *testing.T) {
	testcasesbyitems := map[string][]struct {
		name string
		item *gildedrose.Item
		want *gildedrose.Item
	}{
		"+5 Dexterity Vest Item": {
			{
				name: "when day passes and sell-in field is greater than 0, quality decreases by 1",
				item: &gildedrose.Item{
					Name:    "+5 Dexterity Vest",
					SellIn:  10,
					Quality: 20,
				},
				want: &gildedrose.Item{
					Name:    "+5 Dexterity Vest",
					SellIn:  9,
					Quality: 19,
				},
			},
			{
				name: "when day passes and sell-in field is lower or equal than 0, quality decreases by 2",
				item: &gildedrose.Item{
					Name:    "+5 Dexterity Vest",
					SellIn:  0,
					Quality: 20,
				},
				want: &gildedrose.Item{
					Name:    "+5 Dexterity Vest",
					SellIn:  -1,
					Quality: 18,
				},
			},
		},
		"Aged Brie Item": {
			{
				name: "when day passes and sell-in field is greater than 0, quality increases by 1",
				item: &gildedrose.Item{
					Name:    "Aged Brie",
					SellIn:  2,
					Quality: 0,
				},
				want: &gildedrose.Item{
					Name:    "Aged Brie",
					SellIn:  1,
					Quality: 1,
				},
			},
			{
				name: "when time goes and sell-in field is lower or equal than 0, quality increases by 2",
				item: &gildedrose.Item{
					Name:    "Aged Brie",
					SellIn:  0,
					Quality: 0,
				},
				want: &gildedrose.Item{
					Name:    "Aged Brie",
					SellIn:  -1,
					Quality: 2,
				},
			},
		},
		"Sulfuras, Hand of Ragnaros Item": {
			{
				name: "when day passes quality doesn't change",
				item: &gildedrose.Item{
					Name:    "Sulfuras, Hand of Ragnaros",
					SellIn:  0,
					Quality: 80,
				},
				want: &gildedrose.Item{
					Name:    "Sulfuras, Hand of Ragnaros",
					SellIn:  0,
					Quality: 80,
				},
			},
		},
		"Backstage passes to a TAFKAL80ETC concert Item": {
			{
				name: "when day passes and sell-in field is greater than 10, quality increases by 1",
				item: &gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  15,
					Quality: 20,
				},
				want: &gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  14,
					Quality: 21,
				},
			},
			{
				name: "when day passes and sell-in field is between 10 and 5, quality increases by 2",
				item: &gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  10,
					Quality: 20,
				},
				want: &gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  9,
					Quality: 22,
				},
			},
			{
				name: "when day passes and sell-in field is between 5 and 0, quality increases by 3",
				item: &gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  5,
					Quality: 20,
				},
				want: &gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  4,
					Quality: 23,
				},
			},
			{
				name: "when day passes and sell-in field is equal to 0, quality falls to 0",
				item: &gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  0,
					Quality: 20,
				},
				want: &gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  -1,
					Quality: 0,
				},
			},
		},
	}

	for name, testcases := range testcasesbyitems {
		t.Run(name, func(t *testing.T) {
			for _, testcase := range testcases {
				t.Run(testcase.name, func(t *testing.T) {
					items := []*gildedrose.Item{testcase.item}
					gildedrose.UpdateQuality(items)
					assert.Equal(t, testcase.want.SellIn, testcase.item.SellIn, "sell-in value not expected")
					assert.Equal(t, testcase.want.Quality, testcase.item.Quality, "quality value not expected")
				})
			}
		})
	}
}
