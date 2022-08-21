package catalogue

import (
	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/myugen/go-gildedrose-kata-solution/quality"
	"github.com/myugen/go-gildedrose-kata-solution/sellin"
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
