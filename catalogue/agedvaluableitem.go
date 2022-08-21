package catalogue

import (
	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/myugen/go-gildedrose-kata-solution/quality"
	"github.com/myugen/go-gildedrose-kata-solution/sellin"
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
