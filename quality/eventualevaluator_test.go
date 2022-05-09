package quality_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestEventualEvaluator_Evaluate(t *testing.T) {
	evaluator := quality.NewEventualEvaluator()

	type args struct {
		item gildedrose.Item
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
			name: "should increase by 1 the quality value when an item is provided with sell-in value greater than 10",
			args: args{item: gildedrose.Item{Name: "Eventual-evaluated item", SellIn: 11, Quality: 10}},
			want: want{quality: 11},
		},
		{
			name: "should increase by 2 the quality value when an item is provided with sell-in value lower or equal than 10",
			args: args{item: gildedrose.Item{Name: "Eventual-evaluated item", SellIn: 10, Quality: 11}},
			want: want{quality: 13},
		},
		{
			name: "should increase by 3 the quality value when an item is provided with sell-in value lower or equal than 5",
			args: args{item: gildedrose.Item{Name: "Eventual-evaluated item", SellIn: 5, Quality: 13}},
			want: want{quality: 16},
		},
		{
			name: "should decrease to 0 the quality value when an item is provided with sell-in value lower or equal than 0",
			args: args{item: gildedrose.Item{Name: "Eventual-evaluated item", SellIn: 0, Quality: 16}},
			want: want{quality: 0},
		},
		{
			name: "should keep in 50 the quality value when an item is provided with quality value equal than maximum quality (50) and sell-in value greater than 0",
			args: args{item: gildedrose.Item{Name: "Eventual-evaluated item", SellIn: 4, Quality: quality.RegularMaximumQuality}},
			want: want{quality: quality.RegularMaximumQuality},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			evaluator.Evaluate(&testcase.args.item)
			assert.Equal(t, testcase.want.quality, testcase.args.item.Quality)
		})
	}
}
