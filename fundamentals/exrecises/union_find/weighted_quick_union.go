package unionfind

type weightedQuickUnion struct {
	id    []int
	sizes []int
	count int
}

func NewWeightedQuickUnion(count int) UnionFind {
	id := make([]int, count)
	sizes := make([]int, count)

	for i := range id {
		id[i] = i
	}

	for i := range sizes {
		sizes[i] = 1
	}

	return &weightedQuickUnion{id, sizes, count}
}

func (qf *weightedQuickUnion) Find(site int) int {
	for site != qf.id[site] {
		site = qf.id[site]
	}

	return site
}

func (qf *weightedQuickUnion) Union(aSite, bSite int) {
	aComponent := qf.Find(aSite)
	bComponent := qf.Find(bSite)

	if aComponent == bComponent {
		return
	}

	aSize := qf.sizes[aComponent]
	bSize := qf.sizes[bComponent]

	if aSize > bSize {
		qf.id[bComponent] = aComponent
		qf.sizes[aComponent] += bSize
	} else {
		qf.id[aComponent] = bComponent
		qf.sizes[bComponent] += aSize
	}
	qf.count--
}

func (qf *weightedQuickUnion) Connected(aSite, bSite int) bool {
	return qf.Find(aSite) == qf.Find(bSite)
}

func (qf *weightedQuickUnion) Count() int {
	return qf.count
}
