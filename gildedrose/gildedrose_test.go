package gildedrose_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/stretchr/testify/assert"
)

func TestUpdateQuality(t *testing.T) {
	testcases := map[string]struct {
		item *gildedrose.Item
		want *gildedrose.Item
	}{
		"+5 Dexterity Vest Item": {
			item: &gildedrose.Item{
				Name:    "+5 Dexterity Vest",
				SellIn:  10,
				Quality: 20,
			},
			want: &gildedrose.Item{
				Name:    "+5 Dexterity Vest",
				SellIn:  9,
				Quality: 19,
			},
		},
	}

	for name, testcase := range testcases {
		t.Run(name, func(t *testing.T) {
			items := []*gildedrose.Item{testcase.item}
			gildedrose.UpdateQuality(items)
			assert.Equal(t, testcase.want.SellIn, testcase.item.SellIn)
			assert.Equal(t, testcase.want.Quality, testcase.item.Quality)
		})
	}
}
