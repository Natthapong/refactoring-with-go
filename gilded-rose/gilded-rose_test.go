package main

import (
	"github.com/kr/pretty"
	"reflect"
	"testing"
)

type Cases struct {
	in  []*Item
	out []*Item
}

func TestUpdateQuality(t *testing.T) {
	t.Run("fixture test cases", func(t *testing.T) {
		cases := []Cases{
			{in: []*Item{{"+5 Dexterity Vest", 10, 20}}, out: []*Item{{"+5 Dexterity Vest", 9, 19}}},
			{in: []*Item{{"Aged Brie", 2, 0}}, out: []*Item{{"Aged Brie", 2 - 1, 1}}},
			{in: []*Item{{"Elixir of the Mongoose", 5, 7}}, out: []*Item{{"Elixir of the Mongoose", 5 - 1, 6}}},
			{in: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}, out: []*Item{{"Sulfuras, Hand of Ragnaros", 0, 80}}},
			{in: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}, out: []*Item{{"Sulfuras, Hand of Ragnaros", -1, 80}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15, 1}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 15 - 1, 2}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10, 49}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 10 - 1, 50}}},
			{in: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5, 49}}, out: []*Item{{"Backstage passes to a TAFKAL80ETC concert", 5 - 1, 50}}},
		}

		for _, c := range cases {

			UpdateQuality(c.in)

			if !reflect.DeepEqual(c.in, c.out) {
				t.Error(pretty.Diff(c.in, c.out))
			}
		}
	})
}
