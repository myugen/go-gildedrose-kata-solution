package quality

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
)

// An Evaluator evaluates the amount loss/gain quality
type Evaluator interface {
	Evaluate(item *gildedrose.Item)
}
