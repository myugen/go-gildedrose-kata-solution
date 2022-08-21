package catalogue_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-kata-solution/catalogue"
	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/myugen/go-gildedrose-kata-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestEventualValuableItem_Update(t *testing.T) {
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
			name: "should increase by 1 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value greater than 10",
			args: args{item: item.Item{Name: "Eventual-valuable item", SellIn: 11, Quality: 10}},
			want: want{sellIn: 10, quality: 11},
		},
		{
			name: "should increase by 2 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value lower or equal than 10",
			args: args{item: item.Item{Name: "Eventual-valuable item", SellIn: 10, Quality: 11}},
			want: want{sellIn: 9, quality: 13},
		},
		{
			name: "should increase by 3 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value lower or equal than 5",
			args: args{item: item.Item{Name: "Eventual-valuable item", SellIn: 5, Quality: 13}},
			want: want{sellIn: 4, quality: 16},
		},
		{
			name: "should decrease to 0 the quality value and decrease by 1 the sell-in value when an item is provided with sell-in value lower or equal than 0",
			args: args{item: item.Item{Name: "Eventual-valuable item", SellIn: 0, Quality: 16}},
			want: want{sellIn: -1, quality: 0},
		},
		{
			name: "should keep in 50 the quality value and decrease by 1 the sell-in value when an item is provided with quality value equal than maximum quality (50) and sell-in value greater than 0",
			args: args{item: item.Item{Name: "Eventual-valuable item", SellIn: 4, Quality: quality.RegularMaximumQuality}},
			want: want{sellIn: 3, quality: quality.RegularMaximumQuality},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			valuableItem := catalogue.NewEventualValuableItem(&testcase.args.item)
			valuableItem.Update()

			assert.Equal(t, testcase.want.quality, valuableItem.Item().Quality)
			assert.Equal(t, testcase.want.sellIn, valuableItem.Item().SellIn)
		})
	}
}
