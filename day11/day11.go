package main

import (
	"fmt"
	"math/big"
	"sort"
)

var (
	p = int64(2) * 7 * 13 * 3 * 19 * 17 * 11 * 5

	//p = int64(23) * 19 * 13 * 17
)

func main() {
	monkeys := parse()
	//monkeys := parseTestMonkeys()
	//fmt.Println(monkeys)

	monkeys, inspections := PlayRounds(monkeys, 10000)

	var vals []int
	for _, ins := range inspections {
		vals = append(vals, ins)
	}

	for i, m := range monkeys {
		fmt.Println(i, m.Items)
	}
	fmt.Println(inspections)

	sort.Ints(vals)
	fmt.Println(vals)

	fmt.Println("Max 2:", vals[len(vals)-1]*vals[len(vals)-2])

	//for i, m := range monkeys {
	//	fmt.Println(i, m.Items)
	//}
	//fmt.Println(inspections)
}

func PlayRounds(monkeys []*Monkey, n int) ([]*Monkey, map[int]int) {
	inspectionsByMonkeyIdx := make(map[int]int)

	for i := range monkeys {
		inspectionsByMonkeyIdx[i] = 0
	}

	for i := 0; i < n; i++ {
		fmt.Println("Round", i, monkeys)
		mks, inspections := Round(monkeys)

		for i, val := range inspections {
			inspectionsByMonkeyIdx[i] += val
		}

		monkeys = mks
		//for j, m := range monkeys {
		//	fmt.Println(i, j, m.Items)
		//}

	}

	return monkeys, inspectionsByMonkeyIdx
}

func Round(monkeys []*Monkey) ([]*Monkey, map[int]int) {
	inspectionsByMonkeyIdx := make(map[int]int)

	for i := range monkeys {
		inspectionsByMonkeyIdx[i] = 0
	}

	for i, m := range monkeys {

		inspectionsByMonkeyIdx[i] += len(m.Items)
		throws := m.Turn()

		// reset current items holding
		m.Items = nil

		// distribute thrown items
		for _, throw := range throws {
			monkeys[throw.To].Receive(throw.Item)
		}
	}

	return monkeys, inspectionsByMonkeyIdx
}

func parse() []*Monkey {
	m0 := &Monkey{
		Items: bigInts(84, 66, 62, 69, 88, 91, 91),
		Op: func(old *big.Int) *big.Int {
			old.Mul(old, big.NewInt(11))
			return old
		},
		ThrowTo: func(item *big.Int) int {
			if mod(item, 2) {
				return 4
			}
			return 7
		},
	}

	m1 := &Monkey{
		Items: bigInts(98, 50, 76, 99),
		Op: func(old *big.Int) *big.Int {
			return old.Mul(old, old)
		},
		ThrowTo: func(item *big.Int) int {
			if mod(item, 7) {
				return 3
			}
			return 6
		},
	}
	m2 := &Monkey{
		Items: bigInts(72, 56, 94),
		Op: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(1))
		},
		ThrowTo: func(item *big.Int) int {
			if mod(item, 13) {
				return 4
			}
			return 0
		},
	}

	m3 := &Monkey{
		Items: bigInts(55, 88, 90, 77, 60, 67),
		Op: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(2))
		},
		ThrowTo: func(item *big.Int) int {
			if mod(item, 3) {
				return 6
			}
			return 5
		},
	}

	m4 := &Monkey{
		Items: bigInts(69, 72, 63, 60, 72, 52, 63, 78),
		Op: func(old *big.Int) *big.Int {
			return old.Mul(old, big.NewInt(13))
		},
		ThrowTo: func(item *big.Int) int {
			if mod(item, 19) {
				return 1
			}
			return 7
		},
	}

	m5 := &Monkey{
		Items: bigInts(89, 73),
		Op: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(5))
		},
		ThrowTo: func(item *big.Int) int {
			if mod(item, 17) {
				return 2
			}
			return 0
		},
	}

	m6 := &Monkey{
		Items: bigInts(78, 68, 98, 88, 66),
		Op: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(6))
		},
		ThrowTo: func(item *big.Int) int {
			if mod(item, 11) {
				return 2
			}
			return 5
		},
	}

	m7 := &Monkey{
		Items: bigInts(70),
		Op: func(old *big.Int) *big.Int {
			return old.Add(old, big.NewInt(7))
		},
		ThrowTo: func(item *big.Int) int {
			if mod(item, 5) {
				return 1
			}
			return 3
		},
	}
	return []*Monkey{m0, m1, m2, m3, m4, m5, m6, m7}
}

type Throw struct {
	Item *big.Int
	To   int
}

type Monkey struct {
	Items   []*big.Int
	Op      func(old *big.Int) *big.Int
	ThrowTo func(item *big.Int) int
}

func (m *Monkey) Turn() []Throw {
	var throws []Throw
	for _, item := range m.Items {
		original := item

		worryLevel := m.Op(original)

		// gets bored
		//worryLevel = worryLevel.Div(worryLevel, big.NewInt(3))

		worryLevel = worryLevel.Mod(worryLevel, big.NewInt(p))

		throwTo := m.ThrowTo(worryLevel)
		throws = append(throws, Throw{
			Item: worryLevel,
			To:   throwTo,
		})
	}

	return throws
}

func (m *Monkey) Receive(item *big.Int) {
	m.Items = append(m.Items, item)
}

func filter(items []int, item int) []int {
	var res []int
	for _, i := range items {
		if i != item {
			res = append(res, i)
		}
	}
	return res
}

func parseTestMonkeys() []*Monkey {
	return []*Monkey{
		&Monkey{
			Items: bigInts(79, 98),
			Op: func(old *big.Int) *big.Int {
				return old.Mul(old, big.NewInt(19))
			},
			ThrowTo: func(item *big.Int) int {
				if mod(item, 23) {
					return 2
				}
				return 3
			},
		},
		&Monkey{
			Items: bigInts(54, 65, 75, 74),
			Op: func(old *big.Int) *big.Int {
				old.Add(old, big.NewInt(6))
				return old
			},
			ThrowTo: func(item *big.Int) int {
				if mod(item, 19) {
					return 2
				}
				return 0
			},
		},
		&Monkey{
			Items: bigInts(79, 60, 97),
			Op: func(old *big.Int) *big.Int {
				return old.Mul(old, old)
			},
			ThrowTo: func(item *big.Int) int {
				if mod(item, 13) {
					return 1
				}
				return 3
			},
		},
		&Monkey{
			Items: bigInts(74),
			Op: func(old *big.Int) *big.Int {
				return old.Add(old, big.NewInt(3))
			},
			ThrowTo: func(item *big.Int) int {

				if mod(item, 17) {
					return 0
				}
				return 1
			},
		},
	}
}

func bigInts(nums ...int64) []*big.Int {
	var res []*big.Int
	for _, n := range nums {
		res = append(res, big.NewInt(n))
	}
	return res
}

func mod(a *big.Int, b int) bool {
	n := big.NewInt(0)
	n.Mod(a, big.NewInt(int64(b)))
	if n.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	return false
}
