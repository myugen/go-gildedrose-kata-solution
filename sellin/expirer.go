package sellin

import (
	"github.com/myugen/go-gildedrose-kata-solution/item"
)

// An Expirer expires the sell-in days of an item
type Expirer interface {
	Expire(item *item.Item)
}
