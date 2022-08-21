package sellin

import (
	"github.com/myugen/go-gildedrose-kata-solution/item"
)

type RegularExpirer struct{}

func (RegularExpirer) Expire(item *item.Item) {
	item.SellIn -= 1
}

func NewRegularExpirer() RegularExpirer {
	return RegularExpirer{}
}
