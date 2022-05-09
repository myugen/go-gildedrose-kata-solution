package catalogue

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/myugen/go-gildedrose-solution/sellin"
)

type LegendaryValuableItem struct {
	ValuableItem
}

func NewLegendaryValuableItem(item *gildedrose.Item) LegendaryValuableItem {
	return LegendaryValuableItem{ValuableItem{
		item:      item,
		Evaluator: quality.NewLegendaryItemEvaluator(),
		Expirer:   sellin.NewNoExpirer(),
	}}
}
