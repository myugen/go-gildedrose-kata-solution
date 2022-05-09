package sellin

import (
	"github.com/myugen/go-gildedrose-solution/item"
)

type NoExpirer struct{}

func (NoExpirer) Expire(item *item.Item) {
	item.SellIn += 0
}

func NewNoExpirer() NoExpirer {
	return NoExpirer{}
}
