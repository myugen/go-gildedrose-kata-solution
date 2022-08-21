package catalogue_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-kata-solution/catalogue"
	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/myugen/go-gildedrose-kata-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestLegendaryValuableItem_Update(t *testing.T) {
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
			name: "should keep in 80 the quality and remain the sell-in value when an item is provided",
			args: args{item: item.Item{Name: "Legendary-evaluable item", SellIn: 10, Quality: 9}},
			want: want{sellIn: 10, quality: quality.LegendaryMaximumQuality},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			valuableItem := catalogue.NewLegendaryValuableItem(&testcase.args.item)
			valuableItem.Update()

			assert.Equal(t, testcase.want.quality, valuableItem.Item().Quality)
			assert.Equal(t, testcase.want.sellIn, valuableItem.Item().SellIn)
		})
	}
}
