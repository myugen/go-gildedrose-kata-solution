package catalogue

import (
	"github.com/myugen/go-gildedrose-solution/item"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/myugen/go-gildedrose-solution/sellin"
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
