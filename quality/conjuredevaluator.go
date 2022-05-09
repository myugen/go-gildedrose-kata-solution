package quality

import (
	"github.com/myugen/go-gildedrose-solution/item"
)

type ConjuredEvaluator struct {
	minimumQuality int
}

func (e ConjuredEvaluator) Evaluate(item *item.Item) {
	amountToDecrease := 2
	if item.SellIn <= 0 {
		amountToDecrease = 4
	}
	if item.Quality == e.minimumQuality {
		amountToDecrease = 0
	}

	item.Quality -= amountToDecrease
}

func NewConjuredEvaluator() ConjuredEvaluator {
	return ConjuredEvaluator{minimumQuality: RegularMinimumQuality}
}
