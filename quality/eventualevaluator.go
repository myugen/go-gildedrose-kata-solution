package quality

import (
	"github.com/myugen/go-gildedrose-solution/item"
)

type EventualEvaluator struct {
	maximumQuality int
}

func (e EventualEvaluator) Evaluate(item *item.Item) {
	amountToAdd := 1
	if item.SellIn <= 10 {
		amountToAdd = 2
	}
	if item.SellIn <= 5 {
		amountToAdd = 3
	}
	if item.Quality == e.maximumQuality {
		amountToAdd = 0
	}
	if item.SellIn <= 0 {
		// when sell-in is zero,
		// the quality need to decrease to zero
		amountToAdd = item.Quality * -1
	}

	item.Quality += amountToAdd
}

func NewEventualEvaluator() EventualEvaluator {
	return EventualEvaluator{maximumQuality: RegularMaximumQuality}
}
