package catalogue

import (
	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/myugen/go-gildedrose-kata-solution/quality"
	"github.com/myugen/go-gildedrose-kata-solution/sellin"
)

type ConjuredValuableItem struct {
	ValuableItem
}

func NewConjuredValuableItem(item *item.Item) ConjuredValuableItem {
	return ConjuredValuableItem{ValuableItem{
		item:      item,
		Evaluator: quality.NewConjuredEvaluator(),
		Expirer:   sellin.NewRegularExpirer(),
	}}
}
