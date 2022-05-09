package catalogue

import (
	"github.com/myugen/go-gildedrose-solution/item"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/myugen/go-gildedrose-solution/sellin"
)

type ValuableItem struct {
	item *item.Item
	quality.Evaluator
	sellin.Expirer
}

func (v ValuableItem) Update() {
	v.Evaluate(v.item)
	v.Expire(v.item)
}

func (v ValuableItem) Item() *item.Item {
	return v.item
}
