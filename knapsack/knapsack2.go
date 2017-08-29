package knapsack

import (
	"fmt"
	"math"
	"sort"
)

var Debug = true

// A Packable item is one that can be placed in a Knapsack
// It must implement a Weight() and a Value() function in order to determine
// whether or not the item should be packed or not.
type Packable interface {
	Name() string
	Weight() int64
	Value() int64
	SetName(string)
	SetWeight(int64)
	SetValue(int64)
}

type Knapsackable interface {
	Name() string
	Weight() int64
	Items() []Packable
	AddItem(packable Packable)
	New() Knapsackable
	SetName(string)
	SetWeight(int64)
}

var EPSILON float64 = 0.00000001

func isFloat64Equals(a, b float64) bool {
	if ((a - b) < EPSILON && (b - a) < EPSILON) {
		return true
	}
	return false
}

func greedys(items []Packable, knaps []Knapsackable, i int, c []int64, z *int64, y []int) {
	n := len(items)
	for j := 0; j < n; j++ {
		item := items[j]
		if (y[j] < 0 && item.Weight() <= c[i]) {
			y[j] = i
			c[i] -= item.Weight()
			*z += item.Value()
		}
	}
}

func printPackable(items []Packable) {
	fmt.Printf("items: %d\n", len(items))
	for _, item := range items {
		fmt.Printf("name: %s, weight: %d, value: %d\n", item.Name(), item.Weight(), item.Value())
	}
}

func MultipleKnapsackProblem(items []Packable, knaps []Knapsackable) (packs []Knapsackable, max int64) {
	sort.Slice(knaps, func(i, j int) bool {
		return (knaps[i].Weight() < knaps[j].Weight())
	})
	sort.Slice(items, func(i, j int) bool {
		pwj := float64(items[j].Value()) / float64(items[j].Weight())
		pwi := float64(items[i].Value()) / float64(items[i].Weight())
		if (isFloat64Equals(pwi, pwj)) {
			return items[j].Value() < items[i].Value()
		} else {
			return pwj < pwi
		}
	})

	var z int64
	var y []int

	// 1 initial solution
	n := len(items)
	y = make([]int, n)
	for j := 0; j < n; j++ {
		y[j] = -1
	}

	m := len(knaps)
	c := make([]int64, m)
	for i := 0; i < m; i++ {
		knap := knaps[i]
		c[i] = knap.Weight()
		greedys(items, knaps, i, c, &z, y)
	}

	fmt.Printf("kanpsacks: %d, items: %d\n", m, n)

	if (Debug) {
		fmt.Printf("z1: %d, c: %v, y: %v\n", z, c, y)
	}

	// 2 rearrangement
	z = 0
	for i := 0; i < m; i++ {
		knap := knaps[i]
		c[i] = knap.Weight()
	}

	i := 0
	for j := n - 1; j >= 0; j-- {
		if (y[j] < 0) {
			continue
		}
		item := items[j]
		l := -1
		for k := i; k < m; k++ {
			if (item.Weight() <= c[k]) {
				l = k
				break
			}
		}
		if (l < 0) {
			for k := 0; k < i - 1; k++ {
				if (item.Weight() <= c[k]) {
					l = k
					break
				}
			}
		}
		if (l < 0) {
			y[j] = -1
		} else {
			y[j] = l
			c[l] -= item.Weight()
			z += item.Value()
			if (l < m - 1) {
				i = l + 1
			} else {
				i = 0
			}
		}
	}
	for i := 0; i < m; i++ {
		greedys(items, knaps, i, c, &z, y)
	}
	if (Debug) {
		fmt.Printf("z2: %d, c: %v, y: %v\n", z, c, y)
	}

	// 3 first improvement
	for j := 0; j < n; j++ {
		if (y[j] < 0) {
			continue
		}
		for k := j + 1; k < n; k++ {
			if (y[k] >= 0 && y[k] != y[j]) {
				itemj := items[j]
				itemk := items[k]
				var h int
				var l int
				if (itemj.Weight() >= itemk.Weight()) {
					h = j
					l = k
				} else {
					h = k
					l = j
				}
				d := items[h].Weight() - items[l].Weight()
				u := -1
				var min int64 = math.MaxInt32
				for x := 0; x < n; x++ {
					if (y[x] < 0 && min > items[x].Weight()) {
						min = items[x].Weight()
						u = x
					}
				}
				if (u >= 0 && d <= c[y[l]] && (c[y[h]] + d) >= items[u].Weight()) {
					var p int64
					t := -1
					for x := 0; x < n; x++ {
						if (y[x] < 0 && items[x].Weight() <= (c[y[h]] + d)) {
							if (p < items[x].Value()) {
								p = items[x].Value()
								t = x
							}
						}
					}
					c[y[h]] = c[y[h]] + d - items[t].Weight()
					c[y[l]] = c[y[l]] - d
					y[t] = y[h]
					y[h] = y[l]
					y[l] = y[t]
					z += items[t].Value()
				}
			}
		}
	}
	if (Debug) {
		fmt.Printf("z3: %d, c: %v, y: %v\n", z, c, y)
	}

	// 4 second improvement
	for j := n - 1; j >= 0; j-- {
		if (y[j] < 0) {
			continue
		}
		item := items[j]
		cbar := c[y[j]] + item.Weight()
		var ybig []int
		for k := 0; k < n; k++ {
			if (y[k] < 0 && items[k].Weight() <= cbar) {
				ybig = append(ybig, k)
				cbar = cbar - items[k].Weight()
			}
		}

		var sumpk int64
		for k := 0; k < len(ybig); k++ {
			sumpk += items[ybig[k]].Value()
		}
		if (sumpk > items[j].Value()) {
			for k := 0; k < len(ybig); k++ {
				y[ybig[k]] = y[j]
			}
			c[y[j]] = cbar
			y[j] = -1
			z = z + sumpk - items[j].Value()
		}
	}
	if (Debug) {
		fmt.Printf("z4: %d, c: %v, y: %v\n", z, c, y)
	}

	// 5 pack knapsacks
	max = z
	for i := 0; i < m; i++ {
		knap := knaps[i].New()
		for j := 0; j < n; j++ {
			if (y[j] == i) {
				knap.AddItem(items[j])
			}
		}
		packs = append(packs, knap)
	}
	return
}
