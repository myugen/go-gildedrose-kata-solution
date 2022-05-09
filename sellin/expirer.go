package sellin

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
)

// An Expirer expires the sell-in days of an item
type Expirer interface {
	Expire(item *gildedrose.Item)
}
