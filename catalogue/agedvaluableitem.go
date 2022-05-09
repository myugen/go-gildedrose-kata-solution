package catalogue

import (
	"github.com/myugen/go-gildedrose-solution/item"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/myugen/go-gildedrose-solution/sellin"
)

type AgedValuableItem struct {
	ValuableItem
}

func NewAgedValuableItem(item *item.Item) AgedValuableItem {
	return AgedValuableItem{ValuableItem{
		item:      item,
		Evaluator: quality.NewAgedEvaluator(),
		Expirer:   sellin.NewRegularExpirer(),
	}}
}
