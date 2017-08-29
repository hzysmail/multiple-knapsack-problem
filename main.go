package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"github.com/alexflint/go-arg"
	"github.com/logrusorgru/aurora"
	"io/ioutil"
	"time"
	"math"
	"github.com/hzysmail/multiple-knapsack-problem/knapsack"
)

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
	items []knapsack.Packable
}

func (k *Knap) Weight() int64 {
	return k.weight
}

func (k *Knap) Items() []knapsack.Packable {
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

func (k *Knap) AddItem(packable knapsack.Packable) {
	k.items = append(k.items, packable)
}

func (k *Knap) New() knapsack.Knapsackable {
	kk := Knap{ k.name, k.weight, nil }
	return &kk
}

func readData(fn string) (items []knapsack.Packable, knaps []knapsack.Knapsackable) {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Printf("ERROR: Unable to open file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Reading data from file: %s\n", fn)
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if len(l) > 0 && (l[0] == ';' || l[0] == '#') {
			continue
		}
		fields := strings.Fields(l)
		if len(fields) == 2 {
			num, err := strconv.ParseInt(fields[0], 10, 64)
			if err != nil {
				fmt.Printf("ERROR: Unable to parse value in:\n>>>   %v\n", l)
				continue
			}
			val, err := strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				fmt.Printf("ERROR: Unable to parse value in:\n>>>   %v\n", l)
				continue
			}
			for i := int64(0); i < num; i++ {
				knap := Knap{ name: fmt.Sprintf("knap-%d-%d", i + 1, val), weight: val }
				knaps = append(knaps, &knap)
			}
			continue
		}
		if len(fields) != 3 {
			fmt.Printf("ERROR: Invalid number of fields, must be 3:\n>>>   %v\n", l)
			continue
		}
		val, err := strconv.ParseInt(fields[1], 10,64)
		if err != nil {
			fmt.Printf("ERROR: Unable to parse value in:\n>>>   %v\n", l)
			continue
		}
		weight, err := strconv.ParseInt(fields[2], 10, 64)
		if err != nil {
			fmt.Printf("ERROR: Unable to parse weight in:\n>>>   %v\n", l)
			continue
		}

		item := Item{ name: fields[0], weight: weight, value: val }
		items = append(items, &item)
	}

	// check we got knapsack capacity from input file
	if len(knaps) < 1 {
		fmt.Printf("ERROR: Knapsack empty, probably misformed input file\n")
		os.Exit(1)
	}

	// check we have at least 1 item in store
	if len(items) < 1 {
		fmt.Printf("ERROR: Empty data, probably misformed input file\n")
		os.Exit(1)
	}
	return
}

func mockData() (items []knapsack.Packable, knaps []knapsack.Knapsackable) {
	//items = []knapsack.Packable{
	//	&Item{ "i1", 10, 175 },
	//	&Item{ "i2", 9, 90 },
	//	&Item{ "i3", 4, 20 },
	//	&Item{ "i4", 2, 50 },
	//	&Item{ "i5", 1, 10 },
	//	&Item{ "i6", 20, 200 },
	//}
	//knaps = []knapsack.Knapsackable{
	//	&Knap{ "k1", 20, nil },
	//}

	//items = []knapsack.Packable{
	//	&Item{ "i1", 3, 5 },
	//	&Item{ "i2", 2, 3 },
	//	&Item{ "i3", 1, 4 },
	//}
	//knaps = []knapsack.Knapsackable{
	//	&Knap{ "k1", 5, nil },
	//}

	//items = []knapsack.Packable{
	//	&Item{ "i1", 3, 5 },
	//	&Item{ "i2", 2, 3 },
	//	&Item{ "i3", 1, 4 },
	//	&Item{ "i4", 2, 2 },
	//}
	//knaps = []knapsack.Knapsackable{
	//	&Knap{ "k1", 5, nil },
	//}

	//items = []knapsack.Packable{
	//	&Item{ "i1", 3, 5 },
	//	&Item{ "i2", 3, 3 },
	//}
	//knaps = []knapsack.Knapsackable{
	//	&Knap{ "k1", 5, nil },
	//	&Knap{ "k2", 5, nil },
	//}

	items = []knapsack.Packable{
		&Item{ "i1", 3, 5 },
		&Item{ "i2", 2, 3 },
		&Item{ "i3", 1, 4 },
		&Item{ "i4", 3, 5 },
		&Item{ "i5", 2, 3 },
		&Item{ "i6", 1, 4 },
	}
	knaps = []knapsack.Knapsackable{
		&Knap{ "k1", 5, nil },
		&Knap{ "k2", 5, nil },
	}
	return
}

type Args struct {
	FileName string `arg:"help:输入数据文件名"`
	Output string `arg:"help:输出数据文件名"`
	Divide bool `arg:"help:输出数据是否转成浮点"`
	Power uint `arg:"help:转浮点时除以10的乘方数"`
}

func main() {
	var args Args
	args.FileName = "input"
	args.Output = "output"
	args.Divide = true
	args.Power = 3
	arg.MustParse(&args)

	fmt.Println("start packing...", args)
	defer fmt.Println("end packing...")

	// items, knaps := mockData()
	items, knaps := readData(args.FileName)

	startTime := time.Now().Unix()

	var max int64
	var ret []knapsack.Knapsackable
	// ret, max = knapsack.MultiKnapsack(items, knaps)
	ret, max = knapsack.MultipleKnapsackProblem(items, knaps)
	endTime := time.Now().Unix()

	power := math.Pow(10, float64(args.Power))
	var tmp string
	var data []string
	var num int
	var capacity int64
	for _, knap := range ret {
		num += len(knap.Items())
		capacity += knap.Weight()
	}
	if (args.Divide) {
		tmp = fmt.Sprintf("capacity: %.3f, max: %.3f, num: %d, time: %d seconds",
			float64(capacity) / power, float64(max) / power, num, (endTime - startTime))
	} else {
		tmp = fmt.Sprintf("capacity: %d, max: %d, num: %d, time: %d seconds",
			capacity, max, num, (endTime - startTime))
	}
	fmt.Println(aurora.Red(tmp))
	data = append(data, tmp)
	for _, knap := range ret {
		var tmpItem []string
		var tmpColorItem []string
		var tmpWeight int64
		var tmpNum int
		for _, item := range knap.Items() {
			if (args.Divide) {
				tmp = fmt.Sprintf("    item: name(%s), weight(%.3f), value(%.3f)",
					item.Name(), float64(item.Weight()) / power, float64(item.Value()) / power)
			} else {
				tmp = fmt.Sprintf("    item: name(%s), weight(%d), value(%d)", item.Name(), item.Weight(), item.Value())
			}
			tmpItem = append(tmpItem, tmp)
			tmpColorItem = append(tmpColorItem, aurora.Green(tmp).String())
			tmpWeight += item.Value()
			tmpNum += 1
		}
		if (args.Divide) {
			tmp = fmt.Sprintf("knap: name(%s), capacity(%.3f), used(%.3f), num(%d)",
				knap.Name(), float64(knap.Weight()) / power, float64(tmpWeight) / power, tmpNum)
		} else {
			tmp = fmt.Sprintf("knap: name(%s), capacity(%d), used(%d), num(%d)",
				knap.Name(), knap.Weight(), tmpWeight, tmpNum)
		}
		data = append(data, tmp)
		fmt.Println(aurora.Magenta(tmp))
		data = append(data, tmpItem...)
		for _, item := range tmpColorItem {
			fmt.Println(item)
		}
	}

	writeData(data, args.Output)
}

func writeData(data []string, filename string) {
	ioutil.WriteFile(filename, []byte(strings.Join(data, "\n")), os.ModePerm)
}