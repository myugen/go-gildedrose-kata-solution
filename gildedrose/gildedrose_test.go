package gildedrose_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-kata-solution/catalogue"
	"github.com/myugen/go-gildedrose-kata-solution/gildedrose"
	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/stretchr/testify/assert"
)

func TestUpdateQuality(t *testing.T) {
	type args struct {
		valuableItem catalogue.ValuableItem
	}
	type want struct {
		sellIn  int
		quality int
	}
	testcasesbyitems := map[string][]struct {
		name string
		args args
		want want
	}{
		"+5 Dexterity Vest Item": {
			{
				name: "when day passes and sell-in field is greater than 0, quality decreases by 1",
				args: args{valuableItem: catalogue.NewRegularValuableItem(&item.Item{
					Name:    "+5 Dexterity Vest",
					SellIn:  10,
					Quality: 20,
				}).ValuableItem},
				want: want{
					sellIn:  9,
					quality: 19,
				},
			},
			{
				name: "when day passes and sell-in field is lower or equal than 0, quality decreases by 2",
				args: args{valuableItem: catalogue.NewRegularValuableItem(&item.Item{
					Name:    "+5 Dexterity Vest",
					SellIn:  0,
					Quality: 20,
				}).ValuableItem},
				want: want{
					sellIn:  -1,
					quality: 18,
				},
			},
		},
		"Aged Brie Item": {
			{
				name: "when day passes and sell-in field is greater than 0, quality increases by 1",
				args: args{valuableItem: catalogue.NewAgedValuableItem(&item.Item{
					Name:    "Aged Brie",
					SellIn:  2,
					Quality: 0,
				}).ValuableItem},
				want: want{
					sellIn:  1,
					quality: 1,
				},
			},
			{
				name: "when day passes and sell-in field is lower or equal than 0, quality increases by 2",
				args: args{valuableItem: catalogue.NewAgedValuableItem(&item.Item{
					Name:    "Aged Brie",
					SellIn:  0,
					Quality: 0,
				}).ValuableItem},
				want: want{
					sellIn:  -1,
					quality: 2,
				},
			},
		},
		"Sulfuras, Hand of Ragnaros Item": {
			{
				name: "when day passes quality doesn't change",
				args: args{valuableItem: catalogue.NewLegendaryValuableItem(&item.Item{
					Name:    "Sulfuras, Hand of Ragnaros",
					SellIn:  0,
					Quality: 80,
				}).ValuableItem},
				want: want{
					sellIn:  0,
					quality: 80,
				},
			},
		},
		"Backstage passes to a TAFKAL80ETC concert Item": {
			{
				name: "when day passes and sell-in field is greater than 10, quality increases by 1",
				args: args{valuableItem: catalogue.NewEventualValuableItem(&item.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  15,
					Quality: 20,
				}).ValuableItem},
				want: want{
					sellIn:  14,
					quality: 21,
				},
			},
			{
				name: "when day passes and sell-in field is between 10 and 5, quality increases by 2",
				args: args{valuableItem: catalogue.NewEventualValuableItem(&item.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  10,
					Quality: 20,
				}).ValuableItem},
				want: want{
					sellIn:  9,
					quality: 22,
				},
			},
			{
				name: "when day passes and sell-in field is between 5 and 0, quality increases by 3",
				args: args{valuableItem: catalogue.NewEventualValuableItem(&item.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  5,
					Quality: 20,
				}).ValuableItem},
				want: want{
					sellIn:  4,
					quality: 23,
				},
			},
			{
				name: "when day passes and sell-in field is equal to 0, quality falls to 0",
				args: args{valuableItem: catalogue.NewEventualValuableItem(&item.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  0,
					Quality: 20,
				}).ValuableItem},
				want: want{
					sellIn:  -1,
					quality: 0,
				},
			},
		},
		"Conjured Mana Cake Item": {
			{
				name: "when day passes and sell-in field is greater than 0, quality decreases by 2",
				args: args{valuableItem: catalogue.NewConjuredValuableItem(&item.Item{
					Name:    "Conjured Mana Cake",
					SellIn:  10,
					Quality: 20,
				}).ValuableItem},
				want: want{
					sellIn:  9,
					quality: 18,
				},
			},
			{
				name: "when day passes and sell-in field is lower or equal than 0, quality decreases by 4",
				args: args{valuableItem: catalogue.NewConjuredValuableItem(&item.Item{
					Name:    "Conjured Mana Cake",
					SellIn:  0,
					Quality: 20,
				}).ValuableItem},
				want: want{
					sellIn:  -1,
					quality: 16,
				},
			},
		},
	}

	for name, testcases := range testcasesbyitems {
		t.Run(name, func(t *testing.T) {
			for _, testcase := range testcases {
				t.Run(testcase.name, func(t *testing.T) {
					gildedrose.UpdateQuality(testcase.args.valuableItem)
					assert.Equal(t, testcase.want.sellIn, testcase.args.valuableItem.Item().SellIn, "sell-in value not expected")
					assert.Equal(t, testcase.want.quality, testcase.args.valuableItem.Item().Quality, "quality value not expected")
				})
			}
		})
	}
}
