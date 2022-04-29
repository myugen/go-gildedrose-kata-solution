package gildedrose_test

import (
	"testing"

	"github.com/myugen/go-gildedrose/gildedrose"
)

func TestUpdateQuality(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}
