package catalogue

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/myugen/go-gildedrose-solution/sellin"
)

type EventualValuableItem struct {
	ValuableItem
}

func NewEventualValuableItem(item *gildedrose.Item) EventualValuableItem {
	return EventualValuableItem{ValuableItem{
		item:      item,
		Evaluator: quality.NewEventualEvaluator(),
		Expirer:   sellin.NewRegularExpirer(),
	}}
}
