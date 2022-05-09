package catalogue_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-solution/catalogue"
	"github.com/myugen/go-gildedrose-solution/item"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestRegularValuableItem_Update(t *testing.T) {
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
			name: "should decrease by 1 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value greater than 0",
			args: args{item: item.Item{Name: "Regular-evaluable item", SellIn: 1, Quality: 9}},
			want: want{sellIn: 0, quality: 8},
		},
		{
			name: "should decrease by 2 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value equal than 0",
			args: args{item: item.Item{Name: "Regular-evaluable item", SellIn: 0, Quality: 8}},
			want: want{sellIn: -1, quality: 6},
		},
		{
			name: "should decrease by 2 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value lower than 0",
			args: args{item: item.Item{Name: "Regular-evaluable item", SellIn: -1, Quality: 6}},
			want: want{sellIn: -2, quality: 4},
		},
		{
			name: "should keep in 0 the quality value and decrease by 1 the sell-in value when an item is provided with quality value equal than minimum quality (0)",
			args: args{item: item.Item{Name: "Regular-evaluable item", SellIn: -1, Quality: quality.RegularMinimumQuality}},
			want: want{sellIn: -2, quality: quality.RegularMinimumQuality},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			valuableItem := catalogue.NewRegularValuableItem(&testcase.args.item)
			valuableItem.Update()

			assert.Equal(t, testcase.want.quality, valuableItem.Item().Quality)
			assert.Equal(t, testcase.want.sellIn, valuableItem.Item().SellIn)
		})
	}
}
