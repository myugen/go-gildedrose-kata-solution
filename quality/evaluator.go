package quality

import (
	"github.com/myugen/go-gildedrose-kata-solution/item"
)

// An Evaluator evaluates the amount loss/gain quality
type Evaluator interface {
	Evaluate(item *item.Item)
}
