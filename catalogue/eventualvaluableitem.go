package catalogue

import (
	"github.com/myugen/go-gildedrose-kata-solution/item"
	"github.com/myugen/go-gildedrose-kata-solution/quality"
	"github.com/myugen/go-gildedrose-kata-solution/sellin"
)

type EventualValuableItem struct {
	ValuableItem
}

func NewEventualValuableItem(item *item.Item) EventualValuableItem {
	return EventualValuableItem{ValuableItem{
		item:      item,
		Evaluator: quality.NewEventualEvaluator(),
		Expirer:   sellin.NewRegularExpirer(),
	}}
}
