package sellin_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/myugen/go-gildedrose-solution/sellin"
	"github.com/stretchr/testify/assert"
)

func TestNoExpirer_Expire(t *testing.T) {
	expirer := sellin.NewNoExpirer()

	type args struct {
		item gildedrose.Item
	}
	type want struct {
		sellIn int
	}
	testcases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "should keep the sell-in value when an item is provided",
			args: args{item: gildedrose.Item{Name: "No-expired item", SellIn: 10, Quality: 9}},
			want: want{sellIn: 10},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			expirer.Expire(&testcase.args.item)
			assert.Equal(t, testcase.want.sellIn, testcase.args.item.SellIn, "invalid sell-in value")
		})
	}
}
