package catalogue

import (
	"github.com/myugen/go-gildedrose-solution/item"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/myugen/go-gildedrose-solution/sellin"
)

type RegularValuableItem struct {
	ValuableItem
}

func NewRegularValuableItem(item *item.Item) RegularValuableItem {
	return RegularValuableItem{ValuableItem{
		item:      item,
		Evaluator: quality.NewRegularEvaluator(),
		Expirer:   sellin.NewRegularExpirer(),
	}}
}
