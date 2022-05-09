package quality

import (
	"github.com/myugen/go-gildedrose-solution/item"
)

// A LegendaryEvaluator always evaluates the quality of an item with the maximum legendary quality
type LegendaryEvaluator struct {
	maximumQuality int
}

// Evaluate always returns the legendary quality
func (e LegendaryEvaluator) Evaluate(item *item.Item) {
	item.Quality = e.maximumQuality
}

func NewLegendaryItemEvaluator() LegendaryEvaluator {
	return LegendaryEvaluator{maximumQuality: LegendaryMaximumQuality}
}
