package quality_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-solution/item"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestConjuredEvaluator_Evaluate(t *testing.T) {
	evaluator := quality.NewConjuredEvaluator()

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
			name: "should decrease by 2 the quality value when an item is provided with sell-in value greater than 0",
			args: args{item: item.Item{Name: "Conjured-evaluated item", SellIn: 1, Quality: 10}},
			want: want{quality: 8},
		},
		{
			name: "should decrease by 4 the quality value when an item is provided with sell-in value equal than 0",
			args: args{item: item.Item{Name: "Conjured-evaluated item", SellIn: 0, Quality: 8}},
			want: want{quality: 4},
		},
		{
			name: "should decrease by 4 the quality value when an item is provided with sell-in value lower than 0",
			args: args{item: item.Item{Name: "Regular-evaluated item", SellIn: -1, Quality: 4}},
			want: want{quality: 0},
		},
		{
			name: "should keep the quality value when an item is provided with quality value equal than minimum quality (0)",
			args: args{item: item.Item{Name: "Regular-evaluated item", SellIn: -1, Quality: quality.RegularMinimumQuality}},
			want: want{quality: quality.RegularMinimumQuality},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			evaluator.Evaluate(&testcase.args.item)
			assert.Equal(t, testcase.want.quality, testcase.args.item.Quality)
		})
	}
}
