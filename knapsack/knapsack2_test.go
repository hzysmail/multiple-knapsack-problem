package knapsack

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var items []Packable
var knaps []Knapsackable


type Item struct {
	name string
	weight int64
	value  int64
}

func (i *Item) Weight() int64 {
	return i.weight
}

func (i *Item) Value() int64 {
	return i.value
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) SetName(name string) {
	i.name = name
}

func (i *Item) SetWeight(weight int64) {
	i.weight = weight
}

func (i *Item) SetValue(value int64) {
	i.value = value
}

type Knap struct {
	name string
	weight  int64
	items []Packable
}

func (k *Knap) Weight() int64 {
	return k.weight
}

func (k *Knap) Items() []Packable {
	return k.items
}

func (k *Knap) Name() string {
	return k.name
}

func (k *Knap) SetName(name string) {
	k.name = name
}

func (k *Knap) SetWeight(weight int64) {
	k.weight = weight
}

func (k *Knap) AddItem(packable Packable) {
	k.items = append(k.items, packable)
}

func (k *Knap) New() Knapsackable {
	kk := Knap{ k.name, k.weight, nil }
	return &kk
}


func init() {
	items = []Packable{
		&Item{ "i1", 40, 80 },
		&Item{ "i2", 10, 20 },
		&Item{ "i3", 40, 60 },
		&Item{ "i4", 30, 40 },
		&Item{ "i5", 50, 60 },
		&Item{ "i6", 50, 60 },
		&Item{ "i7", 55, 65 },
		&Item{ "i8", 25, 25 },
		&Item{ "i9", 40, 30 },
	}
	knaps = []Knapsackable{
		&Knap{ "k1", 100, nil },
		&Knap{ "k2", 150, nil },
	}
}

func Test_MultipleKnapsackProblem(t *testing.T) {
	pack, max := MultipleKnapsackProblem(items, knaps)
	assert.Equal(t, int64(350), max)
	assert.Equal(t, 2, len(pack))
	assert.Equal(t, 3, len(pack[0].Items()))
	assert.Equal(t, 4, len(pack[1].Items()))
}
