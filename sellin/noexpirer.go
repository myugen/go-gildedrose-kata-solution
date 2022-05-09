package sellin

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
)

type NoExpirer struct{}

func (NoExpirer) Expire(item *gildedrose.Item) {
	item.SellIn += 0
}

func NewNoExpirer() NoExpirer {
	return NoExpirer{}
}
