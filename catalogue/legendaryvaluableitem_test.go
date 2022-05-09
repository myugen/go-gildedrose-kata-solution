package catalogue_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-solution/catalogue"
	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestLegendaryValuableItem_Update(t *testing.T) {
	type args struct {
		item gildedrose.Item
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
			args: args{item: gildedrose.Item{Name: "Legendary-evaluable item", SellIn: 10, Quality: 9}},
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
