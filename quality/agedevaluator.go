package quality

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
)

type AgedEvaluator struct {
	maximumQuality int
}

func (e AgedEvaluator) Evaluate(item *gildedrose.Item) {
	amountToIncrease := 1
	if item.SellIn <= 0 {
		amountToIncrease = 2
	}
	if item.Quality == e.maximumQuality {
		amountToIncrease = 0
	}

	item.Quality += amountToIncrease
}

func NewAgedEvaluator() AgedEvaluator {
	return AgedEvaluator{maximumQuality: RegularMaximumQuality}
}
