package catalogue

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/myugen/go-gildedrose-solution/sellin"
)

type ValuableItem struct {
	item *gildedrose.Item
	quality.Evaluator
	sellin.Expirer
}

func (v ValuableItem) Update() {
	v.Evaluate(v.item)
	v.Expire(v.item)
}

func (v ValuableItem) Item() *gildedrose.Item {
	return v.item
}
