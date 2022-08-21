package quality_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/myugen/go-gildedrose-kata-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestAgedEvaluator_Evaluate(t *testing.T) {
	evaluator := quality.NewAgedEvaluator()

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
			name: "should increase by 1 the quality value when an item is provided with sell-in value greater than 0",
			args: args{item: item.Item{Name: "Aged-evaluated item", SellIn: 1, Quality: 9}},
			want: want{quality: 10},
		},
		{
			name: "should increase by 2 the quality value when an item is provided with sell-in value equal than 0",
			args: args{item: item.Item{Name: "Aged-evaluated item", SellIn: 0, Quality: 10}},
			want: want{quality: 12},
		},
		{
			name: "should increase by 2 the quality value when an item is provided with sell-in value lower than 0",
			args: args{item: item.Item{Name: "Aged-evaluated item", SellIn: -1, Quality: 12}},
			want: want{quality: 14},
		},
		{
			name: "should keep in 50 the quality value when an item is provided with quality value equal than maximum quality (50)",
			args: args{item: item.Item{Name: "Aged-evaluated item", SellIn: -1, Quality: quality.RegularMaximumQuality}},
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
