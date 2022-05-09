package catalogue

import (
	"github.com/myugen/go-gildedrose-solution/item"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/myugen/go-gildedrose-solution/sellin"
)

type LegendaryValuableItem struct {
	ValuableItem
}

func NewLegendaryValuableItem(item *item.Item) LegendaryValuableItem {
	return LegendaryValuableItem{ValuableItem{
		item:      item,
		Evaluator: quality.NewLegendaryEvaluator(),
		Expirer:   sellin.NewNoExpirer(),
	}}
}
