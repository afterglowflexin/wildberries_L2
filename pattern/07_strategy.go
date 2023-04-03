// Паттерн Strategy определяет набор алгоритмов схожих по роду деятельности,
// инкапсулирует их в отдельный класс и делает их подменяемыми.

package main

type StrategySort interface {
	Sort(a []int)
}

type BubbleSort struct {
}

func (s BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

type InsertionSort struct {
}

func (s InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

type ContextSort struct {
	strategy StrategySort
}

func (c *ContextSort) Set(strategy StrategySort) {
	c.strategy = strategy
}

func (c *ContextSort) Sort(a []int) {
	c.strategy.Sort(a)
}
