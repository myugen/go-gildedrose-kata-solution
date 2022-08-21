package catalogue_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-kata-solution/catalogue"
	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/myugen/go-gildedrose-kata-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestAgedValuableItem_Update(t *testing.T) {
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
			name: "should increase by 1 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value greater than 0",
			args: args{item: item.Item{Name: "Aged-valuable item", SellIn: 1, Quality: 9}},
			want: want{sellIn: 0, quality: 10},
		},
		{
			name: "should increase by 2 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value equal than 0",
			args: args{item: item.Item{Name: "Aged-valuable item", SellIn: 0, Quality: 10}},
			want: want{sellIn: -1, quality: 12},
		},
		{
			name: "should increase by 2 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value lower than 0",
			args: args{item: item.Item{Name: "Aged-valuable item", SellIn: -1, Quality: 12}},
			want: want{sellIn: -2, quality: 14},
		},
		{
			name: "should keep in 50 the quality value and decrease by 1 the sell-in value when an item is provided with quality value equal than maximum quality (50)",
			args: args{item: item.Item{Name: "Aged-valuable item", SellIn: -1, Quality: quality.RegularMaximumQuality}},
			want: want{sellIn: -2, quality: quality.RegularMaximumQuality},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			valuableItem := catalogue.NewAgedValuableItem(&testcase.args.item)
			valuableItem.Update()

			assert.Equal(t, testcase.want.quality, valuableItem.Item().Quality)
			assert.Equal(t, testcase.want.sellIn, valuableItem.Item().SellIn)
		})
	}
}
