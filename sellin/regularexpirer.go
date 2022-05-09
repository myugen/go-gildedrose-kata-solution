package sellin

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
)

type RegularExpirer struct{}

func (RegularExpirer) Expire(item *gildedrose.Item) {
	item.SellIn -= 1
}

func NewRegularExpirer() RegularExpirer {
	return RegularExpirer{}
}
