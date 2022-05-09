package catalogue_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-solution/catalogue"
	"github.com/myugen/go-gildedrose-solution/item"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestConjuredValuableItem_Update(t *testing.T) {
	type args struct {
		item item.Item
	}
	type want struct {
		sellIn  int
		quality int
	}
	testcases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "should decrease by 2 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value greater than 0",
			args: args{item: item.Item{Name: "Conjured-evaluable item", SellIn: 1, Quality: 10}},
			want: want{sellIn: 0, quality: 8},
		},
		{
			name: "should decrease by 4 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value equal than 0",
			args: args{item: item.Item{Name: "Conjured-evaluable item", SellIn: 0, Quality: 8}},
			want: want{sellIn: -1, quality: 4},
		},
		{
			name: "should decrease by 4 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value lower than 0",
			args: args{item: item.Item{Name: "Conjured-evaluable item", SellIn: -1, Quality: 4}},
			want: want{sellIn: -2, quality: 0},
		},
		{
			name: "should keep in 0 the quality value and decrease by 1 the sell-in value when an item is provided with quality value equal than minimum quality (0)",
			args: args{item: item.Item{Name: "Conjured-evaluable item", SellIn: -1, Quality: quality.RegularMinimumQuality}},
			want: want{sellIn: -2, quality: quality.RegularMinimumQuality},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			valuableItem := catalogue.NewConjuredValuableItem(&testcase.args.item)
			valuableItem.Update()

			assert.Equal(t, testcase.want.quality, valuableItem.Item().Quality)
			assert.Equal(t, testcase.want.sellIn, valuableItem.Item().SellIn)
		})
	}
}
