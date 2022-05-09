package catalogue

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/myugen/go-gildedrose-solution/sellin"
)

type AgedValuableItem struct {
	ValuableItem
}

func NewAgedValuableItem(item *gildedrose.Item) AgedValuableItem {
	return AgedValuableItem{ValuableItem{
		item:      item,
		Evaluator: quality.NewAgedEvaluator(),
		Expirer:   sellin.NewRegularExpirer(),
	}}
}
