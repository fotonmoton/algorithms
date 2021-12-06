package sorting

type IntSort []int

func (items IntSort) Len() int {
	return len(items)
}

func (items IntSort) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items IntSort) Less(i, j int) bool {
	return items[i] < items[j]
}
