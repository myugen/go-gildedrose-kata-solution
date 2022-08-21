package quality_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/myugen/go-gildedrose-kata-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestLegendaryEvaluator_Evaluate(t *testing.T) {
	evaluator := quality.NewLegendaryEvaluator()

	type args struct {
		item item.Item
	}
	type want struct {
		quality int
	}
	testcases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "should keep in 80 the quality value when an item is provided",
			args: args{item: item.Item{Name: "Legendary-evaluated item", SellIn: 10, Quality: 9}},
			want: want{quality: quality.LegendaryMaximumQuality},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			evaluator.Evaluate(&testcase.args.item)
			assert.Equal(t, testcase.want.quality, testcase.args.item.Quality)
		})
	}
}
